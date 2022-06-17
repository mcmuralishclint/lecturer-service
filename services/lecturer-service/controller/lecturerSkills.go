package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/middleware"
	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/models"
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

func DeleteLecturerSkills(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	skill, ok := vars["skill"]
	if !ok {
		json.NewEncoder(res).Encode(errors.New("input a valid name map"))
		return
	}
	err := models.DeleteLecturerSkill(skill, middleware.CurrentUserEmail)
	if err != nil {
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	json.NewEncoder(res).Encode(nil)
}
