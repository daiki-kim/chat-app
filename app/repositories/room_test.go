package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/pkg/tester"
	"github.com/stretchr/testify/suite"
)

type RoomeTestSuite struct {
	tester.DBSQLiteSuite
	originalDB *sql.DB
}

func TestRoomTestSuite(t *testing.T) {
	suite.Run(t, new(RoomeTestSuite))
}

func (suite *RoomeTestSuite) SetupSuite() {
	suite.DBSQLiteSuite.SetupSuite()
	suite.originalDB = models.DB
}

func (suite *RoomeTestSuite) TestCreateRoom() {

}
