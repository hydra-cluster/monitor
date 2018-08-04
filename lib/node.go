package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

// Nodes represents a list of registered nodes
type Nodes struct {
	Registered []Node `json:"registered"`
}

// Register node on the registered agents array
func (nodes *Nodes) Register(node Node) {
	nodes.Unregister(node)
	nodes.Registered = append(nodes.Registered, node)
}

// Unregister node on the registered agents array
func (nodes *Nodes) Unregister(node Node) {
	for index, n := range nodes.Registered {
		if n.Hostname == node.Hostname {
			nodes.Registered = append(nodes.Registered[:index], nodes.Registered[index+1:]...)
		}
	}
}

// Save the registered agents to file
func (nodes *Nodes) Save() {
	nodesJSON, _ := json.Marshal(nodes)
	ioutil.WriteFile("registered_agents.json", nodesJSON, os.ModePerm)
}

// Load registered nodes from file
func (nodes *Nodes) Load() {
	nodes.Registered = nil
	jsonByte, _ := ioutil.ReadFile("registered_agents.json")
	json.Unmarshal(jsonByte, nodes)
}

// Node represents one of the Cluster node
type Node struct {
	Hostname           string             `json:"hostname"`
	Distro             string             `json:"distro"`
	Kernel             string             `json:"kernel"`
	Model              string             `json:"model"`
	IP                 string             `json:"ip"`
	Params             []Param            `json:"params"`
	LastConnectionDate time.Time          `json:"last_connection_at"`
	LastUpdatedDate    time.Time          `json:"last_updated_at"`
	NetworkInterfaces  []NetworkInterface `json:"network_interfaces"`
}

// Param represents a instance of servers monitored parameter
type Param struct {
	Label   string `json:"label"`
	Value   string `json:"value"`
	Unit    string `json:"unit"`
	Warning string `json:"warning_target"`
	Danger  string `json:"danger_target"`
}

// NetworkInterface represents server network card
type NetworkInterface struct {
	Name string `json:"name"`
	MAC  string `json:"mac_addr"`
	IP   string `json:"ip"`
}

// Update reload node attributes
func (n *Node) Update() {
	n.LastUpdatedDate = time.Now()
	n.Hostname, _ = os.Hostname()
	n.IP = getOutboundIP().To4().String()
	loadServerParametes(n)
	getInterfaces(&n.NetworkInterfaces)
}

// NewNode returns a new node intance
func NewNode() *Node {
	n := Node{}
	n.LastConnectionDate = time.Now()
	n.Update()
	return &n
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func getInterfaces(nis *[]NetworkInterface) {
	*nis = nil
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		ni := NetworkInterface{}
		ni.Name = i.Name
		ni.MAC = i.HardwareAddr.String()
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ni.IP = ipnet.IP.String()
				}
			}
		}
		if ni.IP != "" {
			*nis = append(*nis, ni)
		}
	}
}

// ExecCommandFolder defines the system path to the lib in python to execute commands
var ExecCommandFolder string

func loadServerParametes(n *Node) {
	out, err := exec.Command("python", ExecCommandFolder+"execCommand.py").Output()
	if err != nil {
		log.Fatalln("lib execCommand.py not found")
	}
	json.Unmarshal(out, n)
}
