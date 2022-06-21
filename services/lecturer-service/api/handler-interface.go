package api

import "net/http"

type SkillsHandler interface {
	Find(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	FindAll(http.ResponseWriter, *http.Request)
}
