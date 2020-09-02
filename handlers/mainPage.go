package handlers

import (
	"html/template"
	"net/http"
)

var mainTpl = template.Must(template.ParseFiles("templates/main.html"))

func MainPage(w http.ResponseWriter, r *http.Request) {
	err := mainTpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "executing template, "+err.Error(), http.StatusInternalServerError)
		return
	}
}
