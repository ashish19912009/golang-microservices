package app

import (
	"net/http"

	"github.com/ashish19912009/golang-microservices/mvc/controllers"
)

// StartApp Package
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
