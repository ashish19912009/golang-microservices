package controllers

import (
	"net/http"
	"strconv"

	"github.com/ashish19912009/golang-microservices/mvc/services"
	"github.com/ashish19912009/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
)

// GetUser Package
func GetUser(c *gin.Context) {
	userID := c.params("user_id")
	usr, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		// just return the bad request to client
		apiErr := &utils.ApplicationError{
			Message:    "User_id must be number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	user, apiErr := services.UserService.GetUser(usr)
	if apiErr != nil {
		// Handle the Error and Return
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	// return the client
	c.JSON(http.StatusOK, user)
}
