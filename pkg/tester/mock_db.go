package tester

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	args = append([]interface{}{query}, args...)
	returnValues := m.Called(args...)
	return returnValues.Get(0).(sql.Result), returnValues.Error(1)
}
