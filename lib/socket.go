package lib

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
)

var upgrader = websocket.Upgrader{}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan interface{}
}

func (c *Client) run() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		select {
		case message := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteJSON(message); err != nil {
				return
			}
		}
	}
}

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool
	// Inbound messages from the clients.
	broadcast chan interface{}
	// Register requests from the clients.
	register chan *Client
	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan interface{}),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				client.send <- message
			}
		}
	}
}

func wsHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan interface{})}
	client.hub.register <- client

	go client.run()
}

// StartWebsocketServer start listening for websocket messages
func StartWebsocketServer(db *DBConn, port string) {
	hub := newHub()
	go hub.run()

	go db.OnChange(DBNodesTable, hub)

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(hub, w, r)
	})
	log.Println("Serving websocket at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
