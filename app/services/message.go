package services

import (
	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/app/repositories"
)

func CreateMessage(message *models.Message) error {
	return repositories.CreateMessage(message)
}

func GetMessagesForRoom(roomID int) ([]*models.Message, error) {
	return repositories.GetMessagesByRoom(roomID)
}
