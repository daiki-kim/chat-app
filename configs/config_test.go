package configs

import (
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestInitEnv(t *testing.T) {
	err := godotenv.Load("/Users/apple/GoProjects/chat-app/.env")
	if err != nil {
		t.Error("Error loading .env file")
	}

	err = LoadEnv()
	assert.Nil(t, err)

	expectedDBPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	assert.Equal(t, os.Getenv("APP_ENV"), Config.Env)
	assert.Equal(t, os.Getenv("DB_HOST"), Config.DBHost)
	assert.Equal(t, expectedDBPort, Config.DBPort)
	assert.Equal(t, os.Getenv("DB_DRIVER"), Config.DBDriver)
	assert.Equal(t, os.Getenv("DB_NAME"), Config.DBName)
	assert.Equal(t, os.Getenv("DB_USER"), Config.DBUser)
	assert.Equal(t, os.Getenv("DB_PASSWORD"), Config.DBPassword)
	assert.Equal(t, os.Getenv("JWT_SECRET"), Config.JwtSecret)
	assert.Equal(t, true, Config.IsDevelopment())
}
