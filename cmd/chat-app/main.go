package main

import (
	"log"
	"net/http"

	"github.com/daiki-kim/chat-app/pkg/api"
	"github.com/daiki-kim/chat-app/pkg/models"
	"github.com/daiki-kim/chat-app/pkg/redis"
	"github.com/daiki-kim/chat-app/pkg/websocket"
)

func main() {
	models.InitDB()
	redis.InitRedis()

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	r := api.NewRouter()
	http.HandleFunc("/ws", websocket.HandleConnections)

	go websocket.HandleMessages()

	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("error starting http server: ", err)
	}
}
