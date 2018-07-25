package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	m "github.com/hydra-cluster/monitor"
)

var dbAddress, dbName string

func main() {
	flag.StringVar(&dbAddress, "dburl", "localhost:28015", "Database address URL")
	flag.StringVar(&dbName, "dbname", m.DBDefaultName, "Database name")

	flag.Parse()

	fmt.Println("------------------------------------")
	fmt.Println("Hydra Cluster Monitor - Agent - v1.0")
	fmt.Println("------------------------------------")

	db := new(m.DBConn)
	db.Connect(dbAddress, dbName)
	defer db.CloseSession()

	node := m.NewNode(db)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("")
		db.CloseSession()
		log.Println("Agent terminated")
		os.Exit(1)
	}()

	log.Println("Agent ready")
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			//db.Insert("logs", nil)
			node.Flush()
			log.Printf("New node log >> %s", node.Hostname)
		}
	}
}
