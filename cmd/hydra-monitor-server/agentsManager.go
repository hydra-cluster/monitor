package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	monitor "github.com/hydra-cluster/monitor/lib"
)

const (
	jsonFilename = "registered_agents.json"
)

// AgentsManager defines an object to manage agents connected to the server
type AgentsManager struct {
	// Registered list of registered agents
	Registered []monitor.Agent `json:"registered"`
}

// Register node on the registered agents array
func (agents *AgentsManager) register(agent monitor.Agent) {
	agents.unregister(agent)
	agents.Registered = append(agents.Registered, agent)
}

// Unregister node on the registered agents array
func (agents *AgentsManager) unregister(agent monitor.Agent) {
	for index, a := range agents.Registered {
		if a.Hostname == agent.Hostname {
			agents.Registered = append(agents.Registered[:index], agents.Registered[index+1:]...)
		}
	}
}

// Save the registered agents to file
func (agents *AgentsManager) save(folder string) {
	agentsJSON, err := json.Marshal(agents)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(folder+jsonFilename, agentsJSON, os.ModePerm)
}

// Load registered router from
func (agents *AgentsManager) load(folder string) {
	agents.Registered = nil
	jsonByte, err := ioutil.ReadFile(folder + jsonFilename)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonByte, agents)
}
