package domain

import (
	"fmt"
	"net/http"

	"github.com/ashish19912009/golang-microservices/mvc/utils"
)

var (
	passengers = map[int64]*Passenger{
		100: {
			ID:       100,
			FullName: "Charu Hashi Rena",
			Age:      28,
			Gender:   "Male",
			TicketNo: "ABC121DC",
		},
	}
	PassengerDao passengerDao
)

type passengerDao struct{}

//GetPassenger Package
func (u *passengerDao) GetPassenger(passID int64) (*Passenger, *utils.ApplicationError) {
	if pass := passengers[passID]; pass != nil {
		return pass, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintln("Passenger was not found", passID),
		StatusCode: http.StatusNotFound,
		Code:       "passenger_not_found",
	}
}
