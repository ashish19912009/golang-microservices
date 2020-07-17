package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ashish19912009/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Ashish Kumar", LastName: "Rena", Email: "ashish1991.2009@gmail.com"},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

// GetUser Package
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("We are acccessing the database")
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintln("Users was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
