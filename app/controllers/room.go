package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/daiki-kim/chat-app/app/services"
	"github.com/daiki-kim/chat-app/pkg/auth"
	"github.com/daiki-kim/chat-app/pkg/logger"
	"go.uber.org/zap"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	ownerID := auth.GetUserIDFromContext(r.Context())

	var requestBody struct {
		Name    string `json:"name"`
		UserIDs []int  `json:"user_ids"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		logger.Warn("failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := services.CreateRoom(requestBody.Name, ownerID, requestBody.UserIDs)
	if err != nil {
		logger.Error("failed to create room", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("room created successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Room created successfully"})
}

func GetRoomsForUser(w http.ResponseWriter, r *http.Request) {
	ownerID := auth.GetUserIDFromContext(r.Context())

	rooms, err := services.GetRoomsForUser(ownerID)
	if err != nil {
		logger.Error("failed to get rooms for user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("rooms retrieved successfully")
	json.NewEncoder(w).Encode(rooms)
}
