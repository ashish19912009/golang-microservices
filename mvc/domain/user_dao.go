package domain

import (
	"fmt"
	"net/http"

	"github.com/ashish19912009/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Ashish Kumar", LastName: "Rena", Email: "ashish1991.2009@gmail.com"},
	}
)

// GetUser Package
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("Users was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
