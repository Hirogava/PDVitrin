package routes

import (
	"net/http"
	"vitrina/db"
	"vitrina/handlers/api"
	"vitrina/handlers/pages"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, manager *db.Manager) {
	Routes(r, manager)
	ApiRoutes(r, manager)
}

func Routes(r *mux.Router, manager *db.Manager) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Projects(w, r, manager)
	}).Methods(http.MethodGet)
}

func ApiRoutes(r *mux.Router, manager *db.Manager) {
	apiPref := r.PathPrefix("/api").Subrouter()

	apiPref.HandleFunc("/project/{id}", func(w http.ResponseWriter, r *http.Request) {
		api.Project(w, r, manager)
	}).Methods(http.MethodGet)

	apiPref.HandleFunc("/filter-projects", func(w http.ResponseWriter, r *http.Request) {
		api.FilterProjects(w, r, manager)
	}).Methods(http.MethodPost)
}
