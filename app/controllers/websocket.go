package controllers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/app/services"
	"github.com/daiki-kim/chat-app/pkg/logger"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn   *websocket.Conn
	RoomID int
	UserID int
}

var clients = make(map[*Client]bool)
var broadcast = make(chan models.Message)
var mutex = &sync.Mutex{}

func handleMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			if client.RoomID == msg.RoomID {
				err := client.Conn.WriteJSON(msg)
				if err != nil {
					client.Conn.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func ChatRoom(w http.ResponseWriter, r *http.Request) {
	roomIDStr := r.URL.Query().Get("room_id")
	roomID, err := strconv.Atoi(roomIDStr)
	userID := jwt.GetUserIDFromContext(r.Context())

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("failed to upgrade connection", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add the client to the list of connected clients
	client := &Client{
		Conn:   conn,
		RoomID: roomID,
		UserID: userID,
	}
	mutex.Lock()
	clients[client] = true
	mutex.Unlock()

	go handleMessages()

	// Listen for messages from the client
	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			mutex.Lock()
			delete(clients, client)
			mutex.Unlock()
			break
		}
		msg.SenderID = userID
		msg.RoomID = roomID

		err = services.CreateMessage(msg.RoomID, msg.SenderID, msg.Content)
		if err != nil {
			logger.Error("failed to create message", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		broadcast <- msg
	}

	// Clean up when the client disconnects
	mutex.Lock()
	delete(clients, client)
	mutex.Unlock()
}
