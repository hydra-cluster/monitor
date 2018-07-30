package lib

import (
	"log"
	"time"

	r "gopkg.in/gorethink/gorethink.v4"
)

const (
	// DBDefaultName defines the default database name
	DBDefaultName string = "hydra_cluster_monitor"
	// DBNodesTable defines the database table with the nodes
	DBNodesTable string = "nodes"
)

// DBConn struct to define the connection to the RethinkDB
type DBConn struct {
	session *r.Session
}

// Connect to the RethinkDB instance
func (db *DBConn) Connect(addr string) {
	log.Printf("Connecting to %s | DB Name: %s", addr, DBDefaultName)
	attemps := 0
	for {
		session, err := r.Connect(r.ConnectOpts{
			Address:  addr,
			Database: DBDefaultName,
		})
		if err != nil {
			attemps++
			time.Sleep(5 * time.Second)
			if attemps > 10 {
				log.Fatalln(err.Error())
			}
			log.Printf("Retrying to connect to DB [%02d]", attemps)
		} else {
			_, err := r.DB(DBDefaultName).Table(DBNodesTable).Run(session)
			if err != nil {
				attemps++
				session.Close()
				time.Sleep(5 * time.Second)
				if attemps > 20 {
					log.Fatalln(err.Error())
				}
				log.Printf("Retrying to connect to DB [%02d]", attemps)
			} else {
				db.session = session
				break
			}
		}
	}
	log.Println("Connected to DB")
}

// Insert a record to the database in the defined table
func (db *DBConn) Insert(tableName string, record interface{}) {
	_, err := r.DB(DBDefaultName).Table(tableName).Insert(record).RunWrite(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// Update a record to the database in the defined table
func (db *DBConn) Update(tableName, id string, record interface{}) {
	_, err := r.DB(DBDefaultName).Table(tableName).Get(id).Update(record).RunWrite(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// DeleteNode unregster a node in the database
func (db *DBConn) DeleteNode(hostname string) {
	table := r.DB(DBDefaultName).Table(DBNodesTable)
	rows, err := table.GetAllByIndex("hostname", hostname).Run(db.session)
	if rows.IsNil() {
		log.Printf("Node %s not registered", hostname)
		return
	}
	n := Node{}
	rows.One(&n)
	err = table.Get(n.ID).Delete().Exec(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("Node %s uregistered", hostname)
}

// LoadNode returns a node registered from the database
func (db *DBConn) LoadNode(nodeRef *Node) {
	query := r.DB(DBDefaultName).Table(DBNodesTable).Filter(r.Row.Field("hostname").Eq(nodeRef.Hostname))
	res, err := query.Run(db.session)
	defer res.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
	if res.IsNil() {
		return
	}
	err = res.One(nodeRef)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// Init creates the database and the initial tables
func (db *DBConn) Init() {
	err := r.DBCreate(DBDefaultName).Exec(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
	rDB := r.DB(DBDefaultName)
	err = rDB.TableCreate(DBNodesTable).Exec(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = rDB.Table(DBNodesTable).IndexCreate("hostname").Exec(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Database initialized")
}

// CloseSession close the current session with the database
func (db *DBConn) CloseSession() {
	db.session.Close()
	log.Println("Database connection closed")
}

// OnChange boradcast message to clients when defined table changes
func (db *DBConn) OnChange(table string, hub *Hub) {
	res, err := r.Table(table).Changes().Run(db.session)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Monitoring changes on table: %s", table)
	var value interface{}
	for res.Next(&value) {
		hub.broadcast <- value
	}
}
