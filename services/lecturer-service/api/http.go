package api

import (
	"encoding/json"
	"net/http"

	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/domain"
)

type handler struct {
	skillService domain.Service
}

func NewHandler(skillService domain.Service) *handler {
	return &handler{skillService: skillService}
}

func (h *handler) Find(w http.ResponseWriter, r *http.Request) {

}
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {

}
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {

}
func (h *handler) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	skills := h.skillService.FindAll()
	json.NewEncoder(w).Encode(skills)
}
