package main

import (
	"net/http"

	"github.com/mcmuralishclint/personal_tutor/lecturer-service/controller"
)

func main() {
	http.HandleFunc("/google/login", controller.GoogleLogin)
	http.HandleFunc("/google/callback", controller.GoogleCallback)
	http.ListenAndServe(":3000", nil)
}
