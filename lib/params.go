package lib

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strconv"
)

const (
	paramDistro       = "distro"
	paramModel        = "model"
	paramKernel       = "kernel"
	paramCPUUsage     = "cpu_usage"
	paramCPUTemp      = "cpu_temp"
	paramRAMUsage     = "ram_usage"
	paramSwapUsage    = "swap_usage"
	paramHDDUsage     = "hdd_usage"
	paramStorageUsage = "storage_usage"
)

// ExecCommandFolder defines the system path to the lib in python to execute commands
var ExecCommandFolder string
var params map[string]interface{}

func loadParam() {
	folder := ExecCommandFolder
	if folder == "" {
		pwd, _ := os.Getwd()
		folder = pwd + "/../../lib/"
	}
	out, _ := exec.Command("python", folder+"execCommand.py").Output()
	var p interface{}
	json.Unmarshal(out, &p)
	if p == nil {
		log.Fatalln("lib execCommand.py not found")
	}
	params = p.(map[string]interface{})
}

func getParam(key string) string {
	if val, ok := params[key]; ok {
		return val.(string)
	}
	return "Undefined"
}

func getFloatParam(key string) float64 {
	s := getParam(key)
	if s == "Undefined" || s == "" {
		return 0
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println(err)
		return 0
	}
	return f
}
