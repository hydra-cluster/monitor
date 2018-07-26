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

var dbAddress, dbName string

func main() {
	var initDB = flag.Bool("dbinit", false, "Create the initial configurations for the database")
	flag.StringVar(&dbAddress, "dburl", "localhost:28015", "Database address URL")
	flag.StringVar(&dbName, "dbname", m.DBDefaultName, "Database name")

	flag.Parse()

	fmt.Println("-------------------------------------")
	fmt.Println("Hydra Cluster Monitor - Server - v1.0")
	fmt.Println("-------------------------------------")

	db := new(m.DBConn)
	db.Connect(dbAddress, dbName)
	defer db.CloseSession()

	if *initDB {
		db.Init()
		return
	}

	go m.StartWebsocketServer(db)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("")
	db.CloseSession()
	log.Println("Server terminated")
	os.Exit(1)
}
