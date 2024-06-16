package services

import (
	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/app/repositories"
)

func CreateRoom(name string, ownerID int, userIDs []int) error {
	room := &models.Room{
		Name:    name,
		OwnerID: ownerID,
	}
	err := repositories.CreateRoom(room)
	if err != nil {
		return err
	}

	roomUserIDs := append(userIDs, ownerID)
	for _, userID := range roomUserIDs {
		err = repositories.AddRoomMember(room.ID, userID)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetRoomsForUser(userID int) ([]*models.Room, error) {
	return repositories.GetRoomsByUser(userID)
}
