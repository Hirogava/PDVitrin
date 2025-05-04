package pages

import (
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

	_, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Неверный номер страницы", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	limit := query.Get("limit")
	if limit == "" {
		limit = "20"
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		log.Println("Неверное количество проектов", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/index.html",
	))
	tmpl.Execute(w, nil)
}