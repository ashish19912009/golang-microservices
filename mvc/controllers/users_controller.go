package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ashish19912009/golang-microservices/mvc/services"
	"github.com/ashish19912009/golang-microservices/mvc/utils"
)

// GetUser Package
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")
	usr, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		// just return the bad request to client
		apiErr := &utils.ApplicationError{
			Message:    "User_id must be number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUser(usr)
	if apiErr != nil {
		// Handle the Error and Return
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	// return the client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
