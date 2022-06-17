package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/models"
)

func Skills(res http.ResponseWriter, req *http.Request) {
	response := make(map[string][]string)
	skills := models.Skills()
	var skill models.Skill
	var allSkills []string

	for skills.Next(context.Background()) {
		if err := skills.Decode(&skill); err != nil {
			log.Fatal(err)
		}
		allSkills = append(allSkills, skill.Value)
	}
	response["skills"] = allSkills
	json.NewEncoder(res).Encode(response)
}

func CreateSkill(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var skill models.Skill
	err := decoder.Decode(&skill)
	if err != nil {
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	fmt.Println(skill)

	err = models.CreateSkill(skill)
	if err != nil {
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	json.NewEncoder(res).Encode(skill)
}

func FindSkill(res http.ResponseWriter, req *http.Request) {
	name_map := req.URL.Query().Get("name_map")
	present, skill := models.FindSkill(name_map)
	if present {
		json.NewEncoder(res).Encode(skill)
		return
	}
	json.NewEncoder(res).Encode(nil)
}

func DeleteSkill(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name_map, ok := vars["name_map"]
	if !ok {
		json.NewEncoder(res).Encode(errors.New("input a valid name map"))
		return
	}
	err := models.DeleteSkill(name_map)
	if err != nil {
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	err = models.DeleteLecturerSkillsBySkill(name_map)
	if err != nil {
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	json.NewEncoder(res).Encode(nil)
}
