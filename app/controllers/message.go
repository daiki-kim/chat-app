package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/daiki-kim/chat-app/app/services"
	"github.com/daiki-kim/chat-app/pkg/auth"
	"github.com/daiki-kim/chat-app/pkg/logger"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	senderID := auth.GetUserIDFromContext(r.Context())
	roomIDStr := mux.Vars(r)["room_id"]
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		logger.Error("failed to parse room id", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestBody struct {
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		logger.Warn("failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.CreateMessage(roomID, senderID, requestBody.Content); err != nil {
		logger.Error("failed to create message", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("message sent successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Message sent successfully"})
}

func GetMessagesForRoom(w http.ResponseWriter, r *http.Request) {
	roomIDStr := mux.Vars(r)["room_id"]
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		logger.Error("failed to parse room id", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messages, err := services.GetMessagesForRoom(roomID)
	if err != nil {
		logger.Error("failed to get messages for room", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("messages retrieved successfully")
	json.NewEncoder(w).Encode(messages)
}
