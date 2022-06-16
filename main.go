package main

import (
	"net/http"

	"github.com/mcmuralishclint/personal_tutor/lecturer-service/controller"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/models"
)

func main() {
	http.HandleFunc("/google/login", controller.GoogleLogin)
	http.HandleFunc("/google/callback", controller.GoogleCallback)
	err := models.ConnectDB()
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":3000", nil)
}
