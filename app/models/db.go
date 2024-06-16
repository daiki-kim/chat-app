package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/daiki-kim/chat-app/configs"
	"github.com/daiki-kim/chat-app/pkg/logger"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const (
	InstanceSqLite int = iota
	InstancePostgres
)

var (
	DB                            *sql.DB
	errInvalidSQLDataBaseInstance = errors.New("invalid sql db instance")
)

func GetModels() []interface{} {
	return []interface{}{&User{}, &Room{}, &Message{}}
}

func NewDatabaseSQLFactory(instance int) (db *sql.DB, err error) {
	switch instance {
	case InstancePostgres:
		dsn := fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			configs.Config.DBUser,
			configs.Config.DBPassword,
			configs.Config.DBHost,
			configs.Config.DBPort,
			configs.Config.DBName,
		)
		db, err = sql.Open("postgres", dsn)
	case InstanceSqLite:
		dsn := fmt.Sprintf("./%s", configs.Config.DBName)
		db, err = sql.Open("sqlite3", dsn)
	default:
		return nil, errInvalidSQLDataBaseInstance
	}
	return db, err
}

func SetDatabase(instance int) (err error) {
	db, err := NewDatabaseSQLFactory(instance)
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
		return err
	}

	if err = db.Ping(); err != nil {
		logger.Fatal("failed to ping database", zap.Error(err))
		return err
	}

	DB = db
	return err
}
