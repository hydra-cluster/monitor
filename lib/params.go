package lib

import (
	"encoding/json"
	"log"
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

var params map[string]interface{}

func loadParam() {
	out, _ := exec.Command("python", "execCommand.py").Output()
	var p interface{}
	json.Unmarshal(out, &p)
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
