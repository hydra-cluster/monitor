package monitor

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v4"
)

const (
	// DBDefaultName defines the default database name
	DBDefaultName string = "hydra_cluster_monitor"
	// DBLogTable defines the database table with the nodes logs
	DBLogTable string = "logs"
	// DBNodesTable defines the database table with the nodes
	DBNodesTable string = "nodes"
)

var databaseName string

// DBConn struct to define the connection to the RethinkDB
type DBConn struct {
	session *r.Session
}

// Connect to the RethinkDB instance
func (db *DBConn) Connect(addr, database string) {
	session, err := r.Connect(r.ConnectOpts{
		Address:  addr,
		Database: database,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	databaseName = database
	db.session = session
	log.Printf("Connected to %s | DB Name: %s", addr, database)
}

// Insert a record to the database in the defined table
func (db *DBConn) Insert(tableName string, record interface{}) {
	_, err := r.DB(databaseName).Table(tableName).Insert(record).RunWrite(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// Update a record to the database in the defined table
func (db *DBConn) Update(tableName, id string, record interface{}) {
	_, err := r.DB(databaseName).Table(tableName).Get(id).Update(record).RunWrite(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// LoadNode returns a node registered from the database
func (db *DBConn) LoadNode(nodeRef *Node) {
	query := r.DB(databaseName).Table(DBNodesTable).Filter(r.Row.Field("hostname").Eq(nodeRef.Hostname))
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
	err := r.DBCreate(databaseName).Exec(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = r.DB(databaseName).TableCreate(DBLogTable).Exec(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = r.DB(databaseName).TableCreate(DBNodesTable).Exec(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = r.DB(databaseName).Table(DBNodesTable).IndexCreate("hostname").Exec(db.session)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Database created")
}

// CloseSession close the current session with the database
func (db *DBConn) CloseSession() {
	db.session.Close()
	log.Println("Database connection closed")
}
