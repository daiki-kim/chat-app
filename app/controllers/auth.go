package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/app/services"
	"github.com/daiki-kim/chat-app/pkg/logger"
	"go.uber.org/zap"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		logger.Warn("failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.RegisterUser(requestBody.Username, requestBody.Email, requestBody.Password)
	if err != nil {
		logger.Error("failed to register user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("user created successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Warn("failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := services.LoginUser(user.Email, user.Password)
	if err != nil {
		logger.Error("failed to login user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("login successful", zap.String("token", token))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
