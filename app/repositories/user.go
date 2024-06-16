package repositories

import "github.com/daiki-kim/chat-app/app/models"

func CreateUser(user *models.User) error {
	_, err := models.DB.Exec(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		user.Username,
		user.Email,
		user.Password,
	)
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	row := models.DB.QueryRow(
		"SELECT id, username, email, password FROM users WHERE email = $1",
		email,
	)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, err
	}
	return user, nil
}
