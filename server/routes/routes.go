package routes

import (
	"net/http"
	"vitrina/db"
	"vitrina/handlers/pages"

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