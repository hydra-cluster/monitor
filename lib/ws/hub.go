package ws

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan interface{}
	webClients chan interface{}
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan interface{}),
		webClients: make(chan interface{}),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

//Emit send message to the hub clients
func (h *Hub) Emit(message *Message) {
	switch message.To {
	case "all":
		h.broadcast <- message
	case "clients":
		h.webClients <- message
	default:
		c := h.getClient(message.To)
		if c != nil {
			c.send <- message
		}
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
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				client.send <- message
			}
		case message := <-h.webClients:
			for client := range h.clients {
				if client.mode == "web" {
					client.send <- message
				}
			}
		}
	}
}

func (h *Hub) getClient(id string) *Client {
	for client := range h.clients {
		if client.id == id {
			return client
		}
	}
	return nil
}
