package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	monitor "github.com/hydra-cluster/monitor/lib"
	socket "github.com/hydra-cluster/monitor/lib/ws"
)

var client *socket.Client
var agent *monitor.Agent

func main() {
	var addr = flag.String("addr", "localhost:5000", "http service address")
	var libFolder = flag.String("lib", "../../lib/", "path to lib execCommand.py folder")

	flag.Parse()

	fmt.Println("--------------------------------------")
	fmt.Println(" Hydra Cluster Monitor - Agent - v1.0 ")
	fmt.Println("--------------------------------------")

	monitor.ExecCommandFolder = *libFolder

	agent = monitor.NewAgent()

	url := "ws://" + *addr + "/ws?id=" + agent.Hostname + "&mode=agent"
	log.Printf("connecting to %s", url)

	client = socket.Dial(url, agent.Hostname, "agent", handlerReadMessage)
	go client.Run()

	log.Println("registering agent")
	client.Emit(socket.NewMessage("server", agent.Hostname, "register_new_agent", "", *agent))

	log.Println("synchronizing Agent")
	time.Sleep(time.Now().Truncate(5 * time.Second).Add(5 * time.Second).Sub(time.Now()))

	log.Println("\033[92magent ready\033[0m")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("")
		client.Close()
		log.Println("Agent terminated")
		os.Exit(1)
	}()

	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			agent.Update()
			agent.Status = "Online"
			err := client.Emit(socket.NewMessage("clients", agent.Hostname, "update_agent_data", "", agent))
			if err != nil {
				return
			}
		}
	}
}

func handlerReadMessage(msg *socket.Message) {
	if msg.Action == "execute_task" {
		task := monitor.Task{}
		jsonBytes, _ := json.Marshal(msg.Content)
		json.Unmarshal(jsonBytes, &task)
		if strings.Contains(task.Target, agent.Hostname) {
			output, err := monitor.ExecuteCommand(task.ParseCommand())
			if err != nil {
				task.Status = "Error"
				task.Output = err.Error()
			} else {
				task.Status = "Done"
				task.Output = string(output)
			}
			client.Emit(socket.NewMessage("clients", agent.Hostname, msg.Action, "200", task))
		}
	}
}
