package services

import (
	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/app/repositories"
)

func CreateRoom(room *models.Room, userIDs []int) error {
	err := repositories.CreateRoom(room)
	if err != nil {
		return err
	}

	for _, userID := range userIDs {
		err = repositories.AddRoomMember(room.ID, userID)
		if err != nil {
			return nil
		}
	}
	return nil
}

func GetRoomsForUser(userID int) ([]*models.Room, error) {
	return repositories.GetRoomsByUser(userID)
}
