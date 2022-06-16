package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/controller"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/middleware"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/models"
)

func main() {
	models.ConnectDB()
	router := mux.NewRouter()

	router.HandleFunc("/google/login", controller.GoogleLogin)
	router.HandleFunc("/google/callback", controller.GoogleCallback)

	skillsRouter := router.PathPrefix("/skills").Subrouter()
	skillsRouter.Use(middleware.IsAuthorized)
	skillsRouter.HandleFunc("/skill/{name_map}", controller.DeleteSkill).Methods("DELETE")
	skillsRouter.HandleFunc("", controller.CreateSkill).Methods("POST")

	router.HandleFunc("/skills", controller.Skills).Methods("GET")
	router.HandleFunc("/skill", controller.FindSkill).Methods("GET")

	http.ListenAndServe(":3000", router)
}
