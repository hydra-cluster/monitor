package lib

import (
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
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
	n.Hostname, _ = os.Hostname()
	n.Model = getServerInfo("dmesg | grep Machine | cut -d ':' -f 4")
	n.Distro = getServerInfo("cat /etc/os-release | head -1 | cut -d '\"' -f 2")
	n.Kernel = getServerInfo("uname -a")
	n.IP = getOutboundIP().To4().String()
	n.LastUpdatedDate = time.Now()
	updateNodeParams(n)
	getInterfaces(&n.NetworkInterfaces)
}

// NewLog returns a instance of Log struct with attributes updated
func updateNodeParams(n *Node) {
	n.Params.CPUUsage = stringToFloat64(getServerInfo("top -bn1 | grep load | awk '{printf \"%.2f\", $(NF-2)}'"))
	n.Params.CPUTemp = stringToFloat64(getServerInfo("vcgencmd measure_temp | cut -d '=' -f 2 | head --bytes -3"))
	n.Params.RAMUsage = stringToFloat64(getServerInfo("free -m | awk 'NR==2{printf \"%.2f\", $3*100/$2 }'"))
	n.Params.SWAPUsage = stringToFloat64(getServerInfo("free -m | awk 'NR==3{printf \"%.2f\", $3*100/$2 }'"))
	n.Params.HDDUsage = stringToFloat64(getServerInfo("df -h | awk '$NF==\"/\"{printf \"%s\", $5}' | head --bytes -1"))
	n.Params.StorageUsage = stringToFloat64(getServerInfo("df -h | awk '$NF==\"/hydra/storage\"{printf \"%s\", $5}' | head --bytes -1"))
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

func getServerInfo(cmd string) string {
	if runtime.GOOS != "linux" {
		return "undefined"
	}
	out, _ := exec.Command(cmd).Output()
	return strings.Trim(string(out), " ")
}

func stringToFloat64(s string) float64 {
	if s == "undefined" {
		return 0
	}
	res, _ := strconv.ParseFloat(s, 64)
	return res
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
