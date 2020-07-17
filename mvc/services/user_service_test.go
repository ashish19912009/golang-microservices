package services

import (
	"net/http"
	"testing"

	"github.com/ashish19912009/golang-microservices/mvc/domain"
	"github.com/ashish19912009/golang-microservices/mvc/utils"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock     usersDaoMock
	getuserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func init() {
	domain.UserDao = &usersDaoMock{}
}
func (m *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getuserFunction(userID)
}
func TestGetUserNotFoundInDatabase(t *testing.T) {
	getuserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 does not exists",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getuserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			ID: 123,
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
}
