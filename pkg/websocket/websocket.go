package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/daiki-kim/chat-app/pkg/db"
	"github.com/daiki-kim/chat-app/pkg/redis"
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
		saveMessage(msg)
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func saveMessage(msg Message) {
	_, err := db.DB.Exec("INSERT INTO messages (username, message) VALUES ($1, $2)", msg.Username, msg.Message)
	if err != nil {
		log.Printf("failed to save message: %v", err)
	}

	msgJSON, err := json.Marshal(msg)
	if err != nil {
		log.Printf("failed to marshal message: %v", err)
		return
	}

	err = redis.Rdb.Set(redis.Ctx, "latestMessage", msgJSON, 0).Err()
	if err != nil {
		log.Printf("failed to cache message: %v", err)
	}
}

func getLatestMessage() (Message, error) {
	var msg Message

	msgJSON, err := redis.Rdb.Get(redis.Ctx, "latestMessage").Result()
	if err != nil {
		return msg, err
	}

	err = json.Unmarshal([]byte(msgJSON), &msg)
	return msg, err
}
