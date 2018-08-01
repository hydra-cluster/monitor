package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	m "github.com/hydra-cluster/monitor/lib"
)

var (
	dbAddress      string
	unregisterNode bool
	libFolder      = "../../lib/"
)

func main() {
	flag.StringVar(&dbAddress, "url", "localhost:28015", "Database address URL")
	flag.BoolVar(&unregisterNode, "unregister", false, "Unregister this node from the database")

	flag.Parse()

	fmt.Println("--------------------------------------")
	fmt.Println(" Hydra Cluster Monitor - Agent - v1.0 ")
	fmt.Println("--------------------------------------")

	db := m.DBConn{}
	db.Connect(dbAddress)
	defer db.CloseSession()

	m.ExecCommandFolder = libFolder

	node := m.NewNode(&db)

	if unregisterNode {
		db.DeleteNode(node.Hostname)
		return
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("")
		db.CloseSession()
		log.Println("Agent terminated")
		os.Exit(1)
	}()

	log.Println("Synchronizing Agent")
	time.Sleep(time.Now().Truncate(5 * time.Second).Add(5 * time.Second).Sub(time.Now()))

	log.Println("Agent ready")
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			node.Update()
			db.Update(m.DBNodesTable, node.ID, node)
		}
	}
}
