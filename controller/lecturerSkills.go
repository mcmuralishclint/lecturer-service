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
	var lectureSkill models.LecturerSkill
	err := decoder.Decode(&lectureSkill)
	if err != nil {
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	// fmt.Println(skill)

	// err = models.CreateSkill(skill)
	// if err != nil {
	// 	json.NewEncoder(res).Encode(err.Error())
	// 	return
	// }
	// json.NewEncoder(res).Encode(skill)
}
