package lib

import (
	"strings"
	"time"
)

const (
	actionShutdown = "Shutdown"
	actionReboot   = "Reboot"
)

// Task represents an command to be executed at the cluster node
type Task struct {
	ID           string        `json:"id"`
	Owner        string        `json:"owner"`
	Action       string        `json:"action"`
	Command      string        `json:"command"`
	Target       string        `json:"target"`
	Status       string        `json:"status"`
	AgentsOutput []AgentOutput `json:"agentsOutput"`
	Start        time.Time     `json:"start"`
	End          time.Time     `json:"end"`
}

// AgentOutput defines the status and output for each agent in the task target
type AgentOutput struct {
	Hostname string    `json:"hostname"`
	Status   string    `json:"status"`
	Output   string    `json:"output"`
	End      time.Time `json:"end"`
}

// ParseCommand validate and parse command before execution
func (t *Task) ParseCommand() string {
	if t.Action == actionShutdown {
		return "sudo shutdown -h +1"
	} else if t.Action == actionReboot {
		return "sudo shutdown -r +1"
	}
	return t.Command
}

//UpdateStatus check if all targets have completed their task and update task status
func (t *Task) UpdateStatus() {
	targets := strings.Split(t.Target, ",")
	if len(targets) == len(t.AgentsOutput) {
		t.End = time.Now()
		for _, o := range t.AgentsOutput {
			if o.Status == "error" {
				t.Status = "Error"
				return
			}
		}
		t.Status = "Done"
	}
}
