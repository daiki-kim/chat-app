package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/app/repositories"
	"github.com/daiki-kim/chat-app/pkg/tester"
	"github.com/stretchr/testify/suite"
)

type MessageTestSuite struct {
	tester.DBSQLiteSuite
	originalDB *sql.DB
}

func TestMessageTestSuite(t *testing.T) {
	suite.Run(t, new(MessageTestSuite))
}

func (suite *MessageTestSuite) SetupSuite() {
	suite.DBSQLiteSuite.SetupSuite()
	suite.originalDB = models.DB
}

func (suite *MessageTestSuite) TestCreateMessage() {
	err := repositories.CreateMessage(&models.Message{
		RoomID:   1,
		SenderID: 1,
		Content:  "test",
	})
	suite.Assert().Nil(err)
}

func (suite *MessageTestSuite) TestGetMessageByRoom() {
	_ = repositories.CreateMessage(&models.Message{
		RoomID:   1,
		SenderID: 1,
		Content:  "test",
	})

	messages, err := repositories.GetMessagesByRoom(1)
	suite.Assert().Nil(err)
	suite.Assert().Equal(1, len(messages))
	suite.Assert().Equal(1, messages[0].RoomID)
	suite.Assert().Equal(1, messages[0].SenderID)
	suite.Assert().NotNil(messages[0].Timestamp)
	suite.Assert().Equal("test", messages[0].Content)
}
