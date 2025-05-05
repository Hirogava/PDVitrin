package pages

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"vitrina/db"
)

func Projects(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	query := r.URL.Query()
	page := query.Get("page")
	if page == "" {
		page = "1"
	}

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Неверный номер страницы: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	limit := query.Get("limit")
	if limit == "" {
		limit = "20"
	}
	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("Неверное количество проектов: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	projectsData, err := manager.GetProjectsMin(pageNumber, limitNumber)
	if err != nil {
		log.Println(fmt.Sprintf("Ошибка получения проектов страницы: %s: ", page),err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	thematicData, err := manager.GetThematic()
	if err != nil {
		log.Println("Ошибка получения тематик: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	specializationData, err := manager.GetSpecializations()
	if err != nil {
		log.Println("Ошибка получения специализаций: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, err := manager.GetProjectsCount()
	if err != nil {
		log.Println("Ошибка получения количества проектов: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	isLastPage := true

	if pageNumber >= count/20 + 1 {
		isLastPage = false
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/index.html",
	))
	tmpl.Execute(w, map[string]interface{}{
		"Projects" : projectsData,
		"Thematic" : thematicData,
		"Specialization" : specializationData,
		"NextPage" : pageNumber + 1,
		"PreviousPage" : pageNumber - 1,
		"IsLastPage" : isLastPage,
		"IsFirstPage" : !(pageNumber == 1),
	})
}