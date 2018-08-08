package ws

import (
	"errors"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var connectionClosed = false

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan interface{}
	// Defines if this client represents an Agent or a Web browser.
	mode string
	// Defines id for hub communication.
	id string
	// Dialer defines if this client is created using a dial to a server.
	dialer bool
	// A function to handler the read message.
	readHandler func(*Hub, *Message)
}

func (c *Client) read() {
	for {
		var msg Message
		if err := c.conn.ReadJSON(&msg); err != nil {
			return
		}
		c.readHandler(c.hub, &msg)
	}
}

func (c *Client) write() {
	defer c.Close()

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

//Close end this client
func (c *Client) Close() {
	c.conn.Close()
	close(c.send)
	connectionClosed = true
	if c.hub != nil {
		c.hub.unregister <- c
	}
	log.Printf("Disconnected %s client: %s", c.mode, c.id)
}

// Run start the client read and write handlers
func (c *Client) Run() {
	go c.read()
	go c.write()
}

// Emit a message to the server
func (c *Client) Emit(msg *Message) error {
	if connectionClosed {
		return errors.New("WebSocket connection closed")
	}
	c.send <- msg
	return nil
}

// Dial creates a new client connected to the server
func Dial(url, id, mode string, readHandler func(*Hub, *Message)) *Client {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	c := &Client{}
	c.conn = conn
	c.send = make(chan interface{})
	c.id = id
	c.mode = mode
	c.dialer = true
	c.readHandler = readHandler
	return c
}
