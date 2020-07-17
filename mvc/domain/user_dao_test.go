package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "We are not expecting user with ID 0")
	assert.NotNil(t, err, "We are expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "Users was not found 0\n", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Ashish Kumar", user.FirstName)
	assert.EqualValues(t, "Rena", user.LastName)
	assert.EqualValues(t, "ashish1991.2009@gmail.com", user.Email)
}

func TestGetPassengerNoUserFound(t *testing.T) {
	user, err := GetPassenger(0)
	assert.Nil(t, user, "We are not expecting user with ID 0")
	assert.NotNil(t, err, "We are expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "passenger_not_found", err.Code)
	assert.EqualValues(t, "Passenger was not found 0\n", err.Message)
}

func TestGetPassengerNoError(t *testing.T) {
	user, err := GetPassenger(100)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 100, user.ID)
	assert.EqualValues(t, "Charu Hashi Rena", user.FullName)
	assert.EqualValues(t, 28, user.Age)
	assert.EqualValues(t, "Male", user.Gender)
}
