package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// StartWebsocketServer start server listening for messages on defined port.
func StartWebsocketServer(port string, readHandler func(*Hub, *Message), registeredNodes interface{}) {
	hub := newHub()
	go hub.run()

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(hub, readHandler, registeredNodes, w, r)
	})
	log.Println("Started WebSocket communications at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func wsHandler(hub *Hub, readHandler func(*Hub, *Message), registeredNodes interface{}, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	val := r.URL.Query()

	client := &Client{
		hub:    hub,
		conn:   conn,
		send:   make(chan interface{}),
		id:     val["id"][0],
		mode:   val["mode"][0],
		dialer: false,
	}
	client.readHandler = readHandler
	go client.Run()

	if client.mode == "web" {
		client.Emit(NewMessage(client.id, "server", "registered-agents", registeredNodes))
	}

	client.hub.register <- client

	log.Printf("Connected %s client: %s", client.mode, client.id)
}
