package lib

import (
	"encoding/json"
	"os"
	"time"
)

// Agent represents each one of the Cluster's node
type Agent struct {
	Hostname           string             `json:"hostname"`
	Status             string             `json:"status"`
	Distro             string             `json:"distro"`
	Kernel             string             `json:"kernel"`
	Model              string             `json:"model"`
	IP                 string             `json:"ip"`
	Params             []param            `json:"params"`
	NetworkInterfaces  []networkInterface `json:"network_interfaces"`
	LastConnectionDate time.Time          `json:"last_connection_at"`
	LastUpdatedDate    time.Time          `json:"last_updated_at"`
}

// Param represents a instance of servers monitored parameter
type param struct {
	Label   string `json:"label"`
	Value   string `json:"value"`
	Unit    string `json:"unit"`
	Warning string `json:"warning"`
	Danger  string `json:"danger"`
}

// NetworkInterface represents server network card
type networkInterface struct {
	Name string `json:"name"`
	MAC  string `json:"mac_addr"`
	IP   string `json:"ip"`
}

// Update reload node attributes
func (agent *Agent) Update() {
	agent.LastUpdatedDate = time.Now()
	agent.Hostname, _ = os.Hostname()
	ip := getOutboundIP()
	if ip == nil {
		agent.IP = "Outbound IP N/D"
	} else {
		agent.IP = ip.To4().String()
	}
	getInterfaces(&agent.NetworkInterfaces)
	outputJSON, _ := ExecuteCommand()
	json.Unmarshal(outputJSON, agent)
}

// NewAgent returns a new node intance
func NewAgent() *Agent {
	agent := Agent{}
	agent.Status = "Offline"
	agent.LastConnectionDate = time.Now()
	agent.Update()
	return &agent
}
