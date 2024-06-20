package main

import (
	"net/http"

	"github.com/daiki-kim/chat-app/app/controllers"
	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/pkg/logger"
	"github.com/daiki-kim/chat-app/pkg/middleware"
	"github.com/gorilla/mux"
)

func main() {
	if err := models.SetDatabase(models.InstancePostgres); err != nil {
		logger.Fatal(err.Error())
	}
	defer models.DB.Close()
	logger.Info("connected to database")

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/signup", controllers.Register).Methods("POST")
	r.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")

	s := r.PathPrefix("/api/v1").Subrouter()
	s.Use(middleware.JwtAuthentication)

	s.HandleFunc("/rooms", controllers.CreateRoom).Methods("POST")
	s.HandleFunc("/rooms", controllers.GetRoomsForUser).Methods("GET")
	s.HandleFunc("/rooms/{room_id}/messages", controllers.CreateMessage).Methods("POST")
	s.HandleFunc("/rooms/{room_id}/messages", controllers.GetMessagesForRoom).Methods("GET")
	s.HandleFunc("/users/{user_id}/rooms", controllers.GetRoomsForUser).Methods("GET")
	s.HandleFunc("/ws/rooms/{room_id}", controllers.ChatRoom).Methods("GET")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	logger.Info("listening on port 8080")
	http.ListenAndServe(":8080", r)
}
