package services

import (
	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/app/repositories"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	return repositories.CreateUser(user)
}

func LoginUser(email, password string) (string, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	// TODO: implement jwt token
	return "token", nil
}
