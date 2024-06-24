package tester

import (
	"os"

	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/configs"
	"github.com/stretchr/testify/suite"
)

type DBSQLiteSuite struct {
	suite.Suite
}

func (suite *DBSQLiteSuite) SetupSuite() {
	configs.Config.DBName = "unittest.sqlite"
	err := models.SetDatabase(models.InstanceSqLite)
	if err != nil {
		suite.Assert().Nil(err)
	}
}

func (suite *DBSQLiteSuite) TearDownSuite() {
	err := os.Remove(configs.Config.DBName)
	suite.Assert().Nil(err)
}
