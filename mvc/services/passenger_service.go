package services

import (
	"github.com/ashish19912009/golang-microservices/mvc/domain"
	"github.com/ashish19912009/golang-microservices/mvc/utils"
)

type passengerService struct{}

var (
	PassengerService passengerService
)

//GetPassenger Package
func (p *passengerService) GetPassenger(passengerID int64) (*domain.Passenger, *utils.ApplicationError) {
	return domain.PassengerDao.GetPassenger(passengerID)
}
