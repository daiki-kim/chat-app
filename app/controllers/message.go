package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/daiki-kim/chat-app/app/services"
	"github.com/daiki-kim/chat-app/pkg/logger"
	"go.uber.org/zap"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	senderID := jwt.GetUserIDFromContext(r.Context())
	roomIDStr := r.URL.Query().Get("room_id")
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Message sent successfully"})
}

func GetMessagesForRoom(w http.ResponseWriter, r *http.Request) {
	roomIDStr := r.URL.Query().Get("room_id")
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

	json.NewEncoder(w).Encode(messages)
}
