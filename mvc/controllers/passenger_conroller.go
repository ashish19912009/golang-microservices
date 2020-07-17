package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ashish19912009/golang-microservices/mvc/services"
	"github.com/ashish19912009/golang-microservices/mvc/utils"
)

// GetPassenger package
func GetPassenger(resp http.ResponseWriter, req *http.Request) {
	passengerID := req.URL.Query().Get("passenger_id")
	pass, err := strconv.ParseInt(passengerID, 10, 64)

	if err != nil {
		urlErr := &utils.ApplicationError{
			Message:    "Passenger Id must be number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(urlErr)
		resp.WriteHeader(urlErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	passenger, apiErr := services.PassengerService.GetPassenger(pass)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	// return to client
	jsonValue, _ := json.Marshal(passenger)
	resp.Write(jsonValue)
}
