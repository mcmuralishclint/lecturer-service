package api

import "net/http"

type SkillsHandler interface {
	FindSkill(http.ResponseWriter, *http.Request)
	CreateSkill(http.ResponseWriter, *http.Request)
	DeleteSkill(http.ResponseWriter, *http.Request)
	FindAllSkills(http.ResponseWriter, *http.Request)
}
