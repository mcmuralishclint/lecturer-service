package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mcmuralishclint/personal_tutor/lecturer-service/middleware"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/models"
)

func AllLecturerSkills(res http.ResponseWriter, req *http.Request) {
	skills := models.AllLecturerSkills(middleware.CurrentUserEmail)
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(skills)
}

func AddLecturerSkills(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var lecturerSkill models.LecturerSkill
	err := decoder.Decode(&lecturerSkill)

	if err != nil {
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	lecturerSkill.Email = middleware.CurrentUserEmail
	_, err = models.FindLecturer(lecturerSkill.Email)
	if err != nil {
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	success, _ := models.FindSkill(lecturerSkill.Skill)
	if success {
		models.AddLecturerSkills(lecturerSkill)
		json.NewEncoder(res).Encode(lecturerSkill)
	}
}
