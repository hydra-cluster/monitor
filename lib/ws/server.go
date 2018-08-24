package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// ServerHub used to interact with the server clients
var ServerHub *Hub
var upgrader = websocket.Upgrader{}

// StartWebsocketServer start server listening for messages on defined port.
func StartWebsocketServer(port string, readHandler func(*Message)) {
	ServerHub = newHub()
	go ServerHub.run()

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(readHandler, w, r)
	})
	log.Println("Started WebSocket communications at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func wsHandler(readHandler func(*Message), w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	val := r.URL.Query()

	client := &Client{
		hub:  ServerHub,
		conn: conn,
		send: make(chan interface{}),
		id:   val["id"][0],
		mode: val["mode"][0],
	}
	client.readHandler = readHandler
	go client.Run()

	client.hub.register <- client
	log.Printf("\033[92mconnected   \033[0m: %s", client.id)
}
