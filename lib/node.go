package lib

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

//Node represents one of the Cluster node
type Node struct {
	ID                 string             `gorethink:"id,omitempty"`
	Hostname           string             `gorethink:"hostname"`
	Distro             string             `gorethink:"distro"`
	Kernel             string             `gorethink:"kernel"`
	Model              string             `gorethink:"model"`
	IP                 string             `gorethink:"ip"`
	Params             []Param            `gorethink:"params"`
	RegisterDate       time.Time          `gorethink:"registered_at"`
	LastConnectionDate time.Time          `gorethink:"last_connection_at"`
	LastUpdatedDate    time.Time          `gorethink:"last_updated_at"`
	NetworkInterfaces  []NetworkInterface `gorethink:"network_interfaces"`
}

// Param represents a instance of servers monitored parameter
type Param struct {
	Label       string         `gorethink:"label"`
	Value       string         `gorethink:"value"`
	Unit        string         `gorethink:"unit"`
	Warning     string         `gorethink:"warning_target"`
	Danger      string         `gorethink:"danger_target"`
	HistoryData []paramHistory `gorethink:"history_data"`
}

func (p *Param) pushToHistory(date time.Time) {
	h := paramHistory{Value: p.Value, Date: date}
	p.HistoryData = append(p.HistoryData, h)
}

type paramHistory struct {
	Value string    `gorethink:"value"`
	Date  time.Time `gorethink:"created_at"`
}

// NetworkInterface represents server network card
type NetworkInterface struct {
	Name string `gorethink:"name"`
	MAC  string `gorethink:"mac_addr"`
	IP   string `gorethink:"ip"`
}

//Update reload node attributes
func (n *Node) Update() {
	n.LastUpdatedDate = time.Now()
	n.Hostname, _ = os.Hostname()
	n.IP = getOutboundIP().To4().String()
	loadServerParametes(n)
	getInterfaces(&n.NetworkInterfaces)
}

// NewNode load node from db or create if it does not exists
func NewNode(db *DBConn) *Node {
	n := Node{}
	n.Hostname, _ = os.Hostname()

	db.LoadNode(&n)

	n.Update()
	n.LastConnectionDate = time.Now()
	if n.ID == "" {
		n.RegisterDate = time.Now()
		db.Insert(DBNodesTable, n)
		log.Printf("Node: %s registered", n.Hostname)
	} else {
		db.Update(DBNodesTable, n.ID, n)
	}
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

	for _, p := range n.Params {
		p.pushToHistory(n.LastUpdatedDate)
	}
}
