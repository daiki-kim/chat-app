package models

type Message struct {
	ID        int    `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Message   string `db:"message" json:"message"`
	Timestamp string `db:"timestamp" json:"timestamp"`
}

func (msg *Message) CreateMessage() error {
	_, err := DB.NamedExec(`INSERT INTO messages (username, message) VALUES (:username, :message)`, msg)
	return err
}

func GetAllMessages() ([]Message, error) {
	var messages []Message
	err := DB.Select(&messages, "SELECT * FROM messages ORDER BY timestamp ASC")
	return messages, err
}

func GetMessageByID(id int) (Message, error) {
	var msg Message
	err := DB.Get(&msg, `SELECT * FROM messages WHERE id = $1`, id)
	return msg, err
}

func (msg *Message) UpdateMessage() error {
	_, err := DB.NamedExec(`UPDATE messages SET message=:message WHERE id=:id`, msg)
	return err
}

func DeleteMessage(id int) error {
	_, err := DB.Exec(`DELETE FROM messages WHERE id = $1`, id)
	return err
}
