package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/daiki-kim/chat-app/pkg/logger"
)

func GetEnvDefault(key, defVal string) string {
	err := godotenv.Load("/Users/apple/GoProjects/chat-app/.env")
	if err != nil {
		logger.Error("Error loading .env file", zap.Error(err))
	}

	val := os.Getenv(key)
	if val == "" {
		return defVal
	}
	return val
}

type ConfigList struct {
	Env                 string
	DBHost              string
	DBPort              int
	DBDriver            string
	DBName              string
	DBUser              string
	DBPassword          string
	APICorsAllowOrigins []string
	JwtSecret           string
}

func (c *ConfigList) IsDevelopment() bool {
	return c.Env == "development"
}

var Config ConfigList

func LoadEnv() error {
	DBPort, err := strconv.Atoi(GetEnvDefault("DB_PORT", "3306"))
	if err != nil {
		return err
	}

	Config = ConfigList{
		Env:                 GetEnvDefault("APP_ENV", "development"),
		DBDriver:            GetEnvDefault("DB_DRIVER", "mysql"),
		DBHost:              GetEnvDefault("DB_HOST", "0.0.0.0"),
		DBPort:              DBPort,
		DBUser:              GetEnvDefault("DB_USER", "app"),
		DBPassword:          GetEnvDefault("DB_PASSWORD", "password"),
		DBName:              GetEnvDefault("DB_NAME", "chatdb"),
		JwtSecret:           GetEnvDefault("JWT_SECRET", "secret"),
		APICorsAllowOrigins: []string{"http://0.0.0.0:8001"},
	}
	return nil
}

func init() {
	if err := LoadEnv(); err != nil {
		logger.Error("Failed to load env: ", zap.Error(err))
	}
}
