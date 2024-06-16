package repositories

import "github.com/daiki-kim/chat-app/app/models"

func CreateMessage(message *models.Message) error {
	_, err := models.DB.Exec(
		"INSERT INTO messages (room_id, sender_id, content) VALUES ($1, $2, $3)",
		message.RoomID,
		message.SenderID,
		message.Content,
	)
	return err
}

func GetMessagesByRoom(roomID int) ([]*models.Message, error) {
	rows, err := models.DB.Query(
		"SELECT id, room_id, sender_id, content, timestamp FROM messages WHERE room_id = $1",
		roomID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := []*models.Message{}
	for rows.Next() {
		message := &models.Message{}
		if err := rows.Scan(&message.ID, &message.SenderID, &message.Content, &message.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
