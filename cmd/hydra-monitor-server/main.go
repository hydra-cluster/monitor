package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hydra-cluster/monitor/lib"
	"github.com/hydra-cluster/monitor/lib/ws"
)

var (
	wsPort          string
	registeredNodes lib.Nodes
	libFolder       = "../../lib/"
)

func main() {
	flag.StringVar(&wsPort, "port", "5000", "WebSocket listening port")
	registeredAgentsJSONFolder := flag.String("f", "", "Path to folder where is going to be saved the JSON file")

	flag.Parse()

	fmt.Println("---------------------------------------")
	fmt.Println(" Hydra Cluster Monitor - Server - v1.0 ")
	fmt.Println("---------------------------------------")

	lib.RegisteredAgentsFolder = *registeredAgentsJSONFolder
	registeredNodes = lib.Nodes{}
	registeredNodes.Load()

	go ws.StartWebsocketServer(wsPort, handlerReadMessage, &registeredNodes)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("")
	log.Println("Server terminated")
	os.Exit(1)
}

func handlerReadMessage(hub *ws.Hub, msg *ws.Message) {
	if msg.To == "server" {
		node := lib.Node{}
		jsonString, _ := json.Marshal(msg.Content)
		json.Unmarshal(jsonString, &node)
		switch msg.Action {
		case "get-registered-agents":
			hub.Emit(ws.NewMessage(msg.From, "server", "registered-agents", registeredNodes))
			return
		case "agent-register":
			registeredNodes.Register(node)
		case "agent-unregister":
			registeredNodes.Unregister(node)
		}
		registeredNodes.Save()
		return
	}
	hub.Emit(msg)
}
