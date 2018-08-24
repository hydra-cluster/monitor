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

const clusterPassword = "pipocadoce"

var (
	agents *AgentsManager
	folder *string
)

func main() {
	port := flag.String("port", "5000", "WebSocket listening port")
	folder = flag.String("f", "", "Path to folder where is going to be saved the JSON file")

	flag.Parse()

	fmt.Println("---------------------------------------")
	fmt.Println(" Hydra Cluster Monitor - Server - v1.0 ")
	fmt.Println("---------------------------------------")

	agents = &AgentsManager{}
	agents.load(*folder)

	go socket.StartWebsocketServer(*port, handlerReadMessage)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("")
	log.Println("Server terminated")
	os.Exit(1)
}

func handlerReadMessage(msg *socket.Message) {
	switch msg.Action {
	case "register_new_agent":
		agent := monitor.Agent{}
		jsonBytes, _ := json.Marshal(msg.Content)
		json.Unmarshal(jsonBytes, &agent)
		agent.Status = "Offline"
		agents.register(agent)
		agents.save(*folder)
	case "registered_agents":
		socket.ServerHub.Emit(
			&socket.Message{
				Action:  msg.Action,
				To:      msg.From,
				From:    "server",
				Status:  "200",
				Content: agents,
			})
	case "execute_task":
		if strings.Contains(msg.From, "client.") {
			var objmap map[string]*json.RawMessage
			jsonBytes, _ := json.Marshal(msg.Content)
			json.Unmarshal(jsonBytes, &objmap)
			var pwd string
			json.Unmarshal(*objmap["password"], &pwd)
			task := monitor.Task{}
			json.Unmarshal(*objmap["task"], &task)
			//invalid password
			if pwd != clusterPassword {
				task.Status = "Invalid password"
				task.End = time.Now()
				socket.ServerHub.Emit(
					&socket.Message{
						Action:  msg.Action,
						To:      msg.From,
						From:    "server",
						Status:  "401",
						Content: task,
					})
				return
			}
			//broadcast to all web clients and agents the new task
			task.Status = "Processing"
			socket.ServerHub.Emit(
				&socket.Message{
					Action:  msg.Action,
					To:      "all",
					From:    "server",
					Status:  "200",
					Content: task,
				})
		} else {
			task := monitor.Task{}
			jsonBytes, _ := json.Marshal(msg.Content)
			json.Unmarshal(jsonBytes, &task)
			task.UpdateStatus()
			msg.Content = task
			socket.ServerHub.Emit(msg)
		}
	default:
		socket.ServerHub.Emit(msg)
	}
}
