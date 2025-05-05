package routes

import (
	"net/http"
	"vitrina/db"
	"vitrina/handlers/pages"
	"vitrina/handlers/api"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, manager *db.Manager) {
	Routes(r, manager)
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
}