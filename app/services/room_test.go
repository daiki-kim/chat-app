package services_test

import (
	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/pkg/tester"
	"github.com/stretchr/testify/suite"
)

type RoomServiceTestSuite struct {
	suite.Suite
	mockDB *tester.MockDB
}

func (suite *RoomServiceTestSuite) SetupSuite() {
	suite.mockDB = new(tester.MockDB)
	models.DB = suite.mockDB
}

func (suite *RoomServiceTestSuite) TearDownSuite() {
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *RoomServiceTestSuite) TestCreateRoom_Success() {
	roomName := "test Room"
	ownerID := 1
	userIDs := []int{2, 3}

	// TODO: Add return room  data and add more test cases
	suite.mockDB.On("Exec", "INSERT INTO rooms (name, owner_id) VALUES ($1, $2) RETURNING id", roomName, ownerID).Return(nil)

}
