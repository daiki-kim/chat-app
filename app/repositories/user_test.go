package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/daiki-kim/chat-app/app/models"
	"github.com/daiki-kim/chat-app/app/repositories"
	"github.com/daiki-kim/chat-app/pkg/tester"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	tester.DBSQLiteSuite
	originalDB *sql.DB
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (suite *UserTestSuite) SetupSuite() {
	suite.DBSQLiteSuite.SetupSuite()
	suite.originalDB = models.DB
}

func (suite *UserTestSuite) TestCreateUser() {
	err := repositories.CreateUser(&models.User{
		Username: "test",
		Email:    "test@example.com",
		Password: "testpass123",
	})
	suite.Assert().Nil(err)
}

func (suite *UserTestSuite) TestGetUserByEmail() {
	_ = repositories.CreateUser(&models.User{
		Username: "test",
		Email:    "test@example.com",
		Password: "testpass123",
	})

	user, err := repositories.GetUserByEmail("test@example.com")
	suite.Assert().Nil(err)
	suite.Assert().Equal("test", user.Username)
	suite.Assert().Equal("testpass123", user.Password)
}
