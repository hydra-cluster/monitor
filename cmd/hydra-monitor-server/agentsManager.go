package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	monitor "github.com/hydra-cluster/monitor/lib"
)

const (
	jsonFilename = "registered_agents.json"
)

type agentsManager []monitor.Agent

// Register node on the registered agents array
func (agents *agentsManager) register(agent monitor.Agent) {
	agents.unregister(agent)
	*agents = append(*agents, agent)
}

// Unregister node on the registered agents array
func (agents *agentsManager) unregister(agent monitor.Agent) {
	for index, a := range *agents {
		if a.Hostname == agent.Hostname {
			*agents = append((*agents)[:index], (*agents)[index+1:]...)
		}
	}
}

// Save the registered agents to file
func (agents *agentsManager) save(folder string) {
	agentsJSON, _ := json.Marshal(agents)
	ioutil.WriteFile(folder+jsonFilename, agentsJSON, os.ModePerm)
}

// Load registered router from
func (agents *agentsManager) load(folder string) {
	agents = nil
	jsonByte, _ := ioutil.ReadFile(folder + jsonFilename)
	json.Unmarshal(jsonByte, agents)
}
