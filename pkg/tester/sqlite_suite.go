package tester

import (
	"database/sql"
	"os"

	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/configs"
	"github.com/daiki-kim/chat-app/pkg/logger"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type DBSQLiteSuite struct {
	suite.Suite
}

func (suite *DBSQLiteSuite) SetupSuite() {
	configs.Config.DBName = "unittest.sqlite"
	err := models.SetDatabase(models.InstanceSqLite)
	suite.Assert().Nil(err)

	err = executeSQLFile(models.DB, "../../external-apps/db/sqlite_init.sql")
	suite.Assert().Nil(err)
}

func (suite *DBSQLiteSuite) TearDownSuite() {
	err := os.Remove(configs.Config.DBName)
	suite.Assert().Nil(err)
}

func executeSQLFile(db *sql.DB, filepath string) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		logger.Error("failed to read sql file", zap.Error(err))
		return err
	}

	_, err = db.Exec(string(file))
	return err
}
