package services

import (
	"github.com/ashish19912009/golang-microservices/mvc/domain"
	"github.com/ashish19912009/golang-microservices/mvc/utils"
)

// GetUser Package
func GetUser(userid int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userid)
}
