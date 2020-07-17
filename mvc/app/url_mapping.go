package app

import (
	"github.com/ashish19912009/golang-microservices/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
	router.GET("/passengers", controllers.GetPassenger)
}
