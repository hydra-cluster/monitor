package lib

import "time"

// Task represents an command to be executed at the cluster node
type Task struct {
	Label   string    `json:"label"`
	Owner   string    `json:"owner"`
	Command string    `json:"command"`
	Target  string    `json:"target"`
	Output  string    `json:"output"`
	Status  string    `json:"status"`
	Start   time.Time `json:"start_at"`
	End     time.Time `json:"end_at"`
}
