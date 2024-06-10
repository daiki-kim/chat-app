package main

import (
	"log"
	"net/http"

	"github.com/daiki-kim/chat-app/pkg/db"
	"github.com/daiki-kim/chat-app/pkg/websocket"
)

func main() {
	db.InitDB()

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", websocket.HandleConnections)

	go websocket.HandleMessages()

	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
	}
}
