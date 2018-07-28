package lib

import (
	"log"
	"net"
	"os"
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
	Params             Params             `gorethink:"params"`
	RegisterDate       time.Time          `gorethink:"registered_at"`
	LastConnectionDate time.Time          `gorethink:"last_connection_at"`
	LastUpdatedDate    time.Time          `gorethink:"last_updated_at"`
	NetworkInterfaces  []NetworkInterface `gorethink:"network_interfaces"`
}

// Params represents attributes for each one of servers monitored params
type Params struct {
	CPUUsage     float64 `gorethink:"cpu_usage"`
	CPUTemp      float64 `gorethink:"cpu_temp"`
	RAMUsage     float64 `gorethink:"ram_usage"`
	SWAPUsage    float64 `gorethink:"swap_usage"`
	HDDUsage     float64 `gorethink:"hdd_usage"`
	StorageUsage float64 `gorethink:"storage_usage"`
}

// NetworkInterface represents server network card
type NetworkInterface struct {
	Name string `gorethink:"name"`
	MAC  string `gorethink:"mac_addr"`
	IP   string `gorethink:"ip"`
}

//Update reload node attributes
func (n *Node) Update() {
	loadParam()
	n.Hostname, _ = os.Hostname()
	n.Model = getParam(paramModel)
	n.Distro = getParam(paramDistro)
	n.Kernel = getParam(paramKernel)
	n.IP = getOutboundIP().To4().String()
	n.LastUpdatedDate = time.Now()
	updateNodeParams(n)
	getInterfaces(&n.NetworkInterfaces)
}

// NewLog returns a instance of Log struct with attributes updated
func updateNodeParams(n *Node) {
	n.Params.CPUUsage = getFloatParam(paramCPUUsage)
	n.Params.CPUTemp = getFloatParam(paramCPUTemp)
	n.Params.RAMUsage = getFloatParam(paramRAMUsage)
	n.Params.SWAPUsage = getFloatParam(paramSwapUsage)
	n.Params.HDDUsage = getFloatParam(paramHDDUsage)
	n.Params.StorageUsage = getFloatParam(paramStorageUsage)
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
