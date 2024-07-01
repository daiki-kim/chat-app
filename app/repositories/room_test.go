package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/app/repositories"
	"github.com/daiki-kim/chat-app/pkg/tester"
	"github.com/stretchr/testify/suite"
)

type RoomTestSuite struct {
	tester.DBSQLiteSuite
	originalDB *sql.DB
}

func AfterTest(suite *RoomTestSuite) {
	models.DB = suite.originalDB
}

func TestRoomTestSuite(t *testing.T) {
	suite.Run(t, new(RoomTestSuite))
}

func (suite *RoomTestSuite) SetupSuite() {
	suite.DBSQLiteSuite.SetupSuite()
	suite.originalDB = models.DB
}

func (suite *RoomTestSuite) TestCreateRoom() {
	room, err := repositories.CreateRoom(&models.Room{
		Name: "test",
	})
	suite.Assert().Nil(err)
	suite.Assert().Equal("test", room.Name)
}

func (suite *RoomTestSuite) TestAddRoomMember() {
	err := repositories.AddRoomMember(1, 2)
	suite.Assert().Nil(err)
}
