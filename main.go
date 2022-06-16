package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/controller"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/models"
)

func main() {
	models.ConnectDB()
	router := mux.NewRouter()

	router.HandleFunc("/google/login", controller.GoogleLogin)
	router.HandleFunc("/google/callback", controller.GoogleCallback)

	router.HandleFunc("/skills", controller.Skills).Methods("GET")
	router.HandleFunc("/skill", controller.FindSkill).Methods("GET")
	router.HandleFunc("/skill/{name_map}", controller.DeleteSkill).Methods("DELETE")
	router.HandleFunc("/skills", controller.CreateSkill).Methods("POST")

	http.ListenAndServe(":3000", router)
}
