package repositories

import "github.com/daiki-kim/chat-app/app/models"

func CreateRoom(room *models.Room) error {
	_, err := models.DB.Exec(
		"INSERT INTO rooms (name, owner_id) VALUES ($1, $2)",
		room.Name,
		room.OwnerID,
	)
	return err
}

func GetRoomsByUser(userID int) ([]*models.Room, error) {
	rows, err := models.DB.Query(
		"SELECT r.id, r.name, r.created_at, r.updated_at, r.owner_id FROM rooms r JOIN room_members rm ON r.id = rm.room_id WHERE rm.user_id = $1",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []*models.Room{}
	for rows.Next() {
		room := &models.Room{}
		if err := rows.Scan(&room.ID, &room.Name, &room.CreatedAt, &room.UpdatedAt, &room.OwnerID); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func AddRoomMember(roomID, userID int) error {
	_, err := models.DB.Exec(
		"INSERT INTO room_members (room_id, user_id) VALUES ($1, $2)",
		roomID,
		userID,
	)
	return err
}
