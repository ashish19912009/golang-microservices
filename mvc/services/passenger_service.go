package services

import (
	"github.com/ashish19912009/golang-microservices/mvc/domain"
	"github.com/ashish19912009/golang-microservices/mvc/utils"
)

//GetPassenger Package
func GetPassenger(passengerID int64) (*domain.Passenger, *utils.ApplicationError) {
	return domain.GetPassenger(passengerID)
}
