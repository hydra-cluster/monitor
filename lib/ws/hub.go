package ws

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// Registered web clients.
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

//Emit send message to the hub clients
func (h *Hub) Emit(message *Message) {
	if message.To == "broadcast" {
		h.broadcast <- message
		return
	}
	h.getClient(message.To).send <- message
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
