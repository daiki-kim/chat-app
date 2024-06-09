package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, conn)
			break
		}
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("erro: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
