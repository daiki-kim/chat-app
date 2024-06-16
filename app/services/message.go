package services

import (
	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/app/repositories"
)

func CreateMessage(roomID, senderID int, content string) error {
	var message = &models.Message{
		RoomID:   roomID,
		SenderID: senderID,
		Content:  content,
	}
	return repositories.CreateMessage(message)
}

func GetMessagesForRoom(roomID int) ([]*models.Message, error) {
	return repositories.GetMessagesByRoom(roomID)
}
