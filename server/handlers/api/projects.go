package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"vitrina/db"
	"vitrina/models"

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

func FilterProjects(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	var filters models.ProjectFilter
	if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
		log.Println("Ошибка декодирования запроса: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	projects, err := manager.GetProjectsByFilters(&filters)
	if err != nil {
		log.Println("Ошибка получения проектов: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}