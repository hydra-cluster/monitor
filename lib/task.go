package lib

import "time"

const (
	actionShutdown = "Shutdown"
	actionReboot   = "Reboot"
)

// Task represents an command to be executed at the cluster node
type Task struct {
	ID      string    `json:"id"`
	Owner   string    `json:"owner"`
	Action  string    `json:"action"`
	Command string    `json:"command"`
	Target  string    `json:"target"`
	Output  string    `json:"output"`
	Status  string    `json:"status"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
}

// ParseCommand validate and parse command before execution
func (t *Task) ParseCommand() string {
	if t.Action == actionShutdown {
		return "sudo shutdown -h"
	} else if t.Action == actionReboot {
		return "sudo reboot"
	}
	return t.Command
}
