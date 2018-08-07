package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	monitor "github.com/hydra-cluster/monitor/lib"
	"github.com/hydra-cluster/monitor/lib/ws"
)

func main() {
	var addr = flag.String("addr", "localhost:5000", "http service address")
	var folder = flag.String("lib", "../../lib/", "path to lib execCommand.py folder")

	flag.Parse()

	fmt.Println("--------------------------------------")
	fmt.Println(" Hydra Cluster Monitor - Agent - v1.0 ")
	fmt.Println("--------------------------------------")

	monitor.ExecCommandFolder = *folder

	node := monitor.NewNode()

	url := "ws://" + *addr + "/ws?id=" + node.Hostname + "&mode=agent"
	log.Printf("connecting to %s", url)

	client := ws.Dial(url, node.Hostname, "agent", handlerReadMessage)
	go client.Run()

	log.Println("Synchronizing Agent")
	time.Sleep(time.Now().Truncate(5 * time.Second).Add(5 * time.Second).Sub(time.Now()))

	log.Println("Registering agent")
	client.Emit(ws.NewMessage("server", node.Hostname, "agent-register", *node))

	log.Println("Agent ready")

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
			node.Update()
			err := client.Emit(ws.NewMessage("broadcast", node.Hostname, "updated-agent", *node))
			if err != nil {
				return
			}
		}
	}
}

func handlerReadMessage(hub *ws.Hub, msg *ws.Message) {

}
