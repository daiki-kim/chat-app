package services

import (
	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/app/repositories"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
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
