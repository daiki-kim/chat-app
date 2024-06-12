package models

import "errors"

type Message struct {
	ID        int    `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Message   string `db:"message" json:"message"`
	Timestamp string `db:"timestamp" json:"timestamp"`
}

func CreateMessage(msg *Message) error {
	_, err := DB.NamedExec(`INSERT INTO messages (username, message) VALUES (:username, :message)`, msg)
	return err
}

func GetAllMessages() ([]Message, error) {
	var messages []Message
	err := DB.Select(&messages, "SELECT * FROM messages ORDER BY timestamp ASC")
	if messages == nil {
		return nil, errors.New("no messages found")
	}
	return messages, err
}

func GetMessageByID(id int) (Message, error) {
	var msg Message
	err := DB.Get(&msg, `SELECT * FROM messages WHERE id = $1`, id)
	if msg == (Message{}) {
		return msg, errors.New("message not found")
	}
	return msg, err
}

func UpdateMessage(msg *Message) error {
	_, err := DB.NamedExec(`UPDATE messages SET message=:message WHERE id=:id`, msg)
	return err
}

func DeleteMessage(id int) error {
	_, err := DB.Exec(`DELETE FROM messages WHERE id = $1`, id)
	return err
}
