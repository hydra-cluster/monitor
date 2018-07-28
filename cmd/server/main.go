package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	m "github.com/hydra-cluster/monitor/lib"
)

var (
	dbInit         bool
	dbAddress      string
	unregisterNode string
	libFolder      = "../../lib/"
)

func main() {
	flag.BoolVar(&dbInit, "init", false, "Create the initial configurations for the database")
	flag.StringVar(&unregisterNode, "unregister", "", "Delete node from the database passing the hostname")
	flag.StringVar(&dbAddress, "url", "localhost:28015", "Database address URL")

	flag.Parse()

	fmt.Println("---------------------------------------")
	fmt.Println(" Hydra Cluster Monitor - Server - v1.0 ")
	fmt.Println("---------------------------------------")

	db := new(m.DBConn)
	db.Connect(dbAddress)
	defer db.CloseSession()

	if dbInit {
		db.Init()
		return
	}

	if unregisterNode != "" {
		db.DeleteNode(unregisterNode)
		return
	}

	go m.StartWebsocketServer(db, "5000")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("")
	db.CloseSession()
	log.Println("Server terminated")
	os.Exit(1)
}
