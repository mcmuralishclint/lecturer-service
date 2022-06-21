package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/domain"
)

type handler struct {
	skillService domain.Service
}

func NewHandler(skillService domain.Service) *handler {
	return &handler{skillService: skillService}
}

func (h *handler) FindSkill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name_map := chi.URLParam(r, "name_map")
	skill := h.skillService.FindSkill(name_map)
	json.NewEncoder(w).Encode(skill)
}
func (h *handler) CreateSkill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var skill domain.Skill
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	existingSkill := h.skillService.FindSkill(skill.NameMap)
	if existingSkill.NameMap != "" {
		json.NewEncoder(w).Encode(errors.New("Record already exists").Error())
		return
	}
	skill, err = h.skillService.CreateSkill(skill)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(skill)
}
func (h *handler) DeleteSkill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name_map := chi.URLParam(r, "name_map")
	success, err := h.skillService.DeleteSkill(name_map)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(success)
}
func (h *handler) FindAllSkills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skills := h.skillService.FindAllSkills()
	json.NewEncoder(w).Encode(skills)
}
