package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/daiki-kim/chat-app/pkg/models"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	err = models.CreateMessage(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := models.GetAllMessages()
	if err != nil {
		if err.Error() == "no messages found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

func GetMessageByID(w http.ResponseWriter, r *http.Request) {
	msgID := r.URL.Query().Get("id")
	if msgID == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	intMsgID, _ := strconv.Atoi(msgID)
	msg, err := models.GetMessageByID(intMsgID)
	if err != nil {
		if err.Error() == "message not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	msgID := r.URL.Query().Get("id")
	if msgID == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	intMsgID, _ := strconv.Atoi(msgID)

	var msg models.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	msg.ID = intMsgID
	err = models.UpdateMessage(&msg)
	if err != nil {
		if err.Error() == "message not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}
