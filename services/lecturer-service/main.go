package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/api"
	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/config"
	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/domain"
	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/repository/mongo"
)

func main() {
	conf, _ := config.NewConfig("/home/muralishc/Downloads/myTutor/myTutor/services/lecturer-service/config/config.yaml")
	repo, _ := mongo.NewMongoRepository(conf.Database.URL, conf.Database.DB, conf.Database.Timeout)
	service := domain.NewSkillService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/skills", handler.FindAllSkills)
	r.Get("/skills/{name_map}", handler.FindSkill)
	r.Delete("/skills/{name_map}", handler.DeleteSkill)
	r.Post("/skills", handler.CreateSkill)
	log.Fatal(http.ListenAndServe(":3000", r))
}
