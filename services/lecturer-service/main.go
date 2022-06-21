package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/api"
	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/domain"
	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/repository"
)

func main() {
	mongoUri := "mongodb+srv://" + "mcmuralishclint" + ":" + "mc159357555" + "@my-personal-professor-v.k20xc.mongodb.net/?retryWrites=true&w=majority"
	repo, _ := repository.NewMongoRepository(mongoUri, "lecturer", 5)
	service := domain.NewSkillService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/skills", handler.FindAll)
	log.Fatal(http.ListenAndServe(":3000", r))
}
