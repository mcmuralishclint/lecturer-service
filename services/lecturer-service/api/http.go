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

func (h *handler) Find(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name_map := chi.URLParam(r, "name_map")
	skill := h.skillService.Find(name_map)
	json.NewEncoder(w).Encode(skill)
}
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var skill domain.Skill
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	existingSkill := h.skillService.Find(skill.NameMap)
	if existingSkill.NameMap != "" {
		json.NewEncoder(w).Encode(errors.New("Record already exists").Error())
		return
	}
	skill, err = h.skillService.Create(skill)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(skill)
}
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name_map := chi.URLParam(r, "name_map")
	success, err := h.skillService.Delete(name_map)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(success)
}
func (h *handler) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skills := h.skillService.FindAll()
	json.NewEncoder(w).Encode(skills)
}
