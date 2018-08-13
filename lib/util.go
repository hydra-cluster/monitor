package lib

import (
	"log"
	"net"
	"os/exec"
)

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func getInterfaces(nis *[]networkInterface) {
	*nis = nil
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		ni := networkInterface{}
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

// ExecuteCommand executes the command in the agent using a python lib
func ExecuteCommand(command ...string) ([]byte, error) {
	if len(command) > 0 {
		return exec.Command("python", ExecCommandFolder+"execCommand.py", "-cmd", command[0]).Output()
	}
	return exec.Command("python", ExecCommandFolder+"execCommand.py").Output()
}
