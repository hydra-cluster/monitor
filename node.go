package monitor

import (
	"log"
	"os"
	"time"
)

//Node represents one of the Cluster node
type Node struct {
	ID                 string    `gorethink:"id,omitempty"`
	Hostname           string    `gorethink:"hostname"`
	RegisterDate       time.Time `gorethink:"registered_at"`
	LastConnectionDate time.Time `gorethink:"last_conn_at"`
}

// Flush update node informations
func (n *Node) Flush() {
	n.Hostname, _ = os.Hostname()
}

// NewNode load node from db or create if it does not exists
func NewNode(db *DBConn) *Node {
	n := Node{}
	n.Hostname, _ = os.Hostname()

	db.LoadNode(&n)

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
