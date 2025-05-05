package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"vitrina/db"

	"github.com/gorilla/mux"
)

func Project(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Неверный id проекта: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	project, err := manager.GetProject(id)
	if err != nil {
		log.Println("Ошибка получения проекта: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}