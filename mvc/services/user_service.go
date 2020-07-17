package services

import (
	"github.com/ashish19912009/golang-microservices/mvc/domain"
	"github.com/ashish19912009/golang-microservices/mvc/utils"
)

type userService struct{}

var (
	UserService userService
)

// GetUser Package
func (u *userService) GetUser(userid int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
