package handlers

import (
	"html/template"
	"net/http"
)

type createPage struct {
	OriginalURL string
	TinyUrl     string
}

type UrlCreator interface {
	CreateURL(host string, inputUrl string) (string, error)
}

type creatingURL struct {
	service UrlCreator
	tpl     *template.Template
}

func NewCreatingURL(service UrlCreator) *creatingURL {
	return &creatingURL{
		service: service,
		tpl:     template.Must(template.ParseFiles("templates/create.html")),
	}
}

func (h *creatingURL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	inputUrl := r.FormValue("inputUrl")

	tinyURL, err := h.service.CreateURL(r.Host, inputUrl)
	if err != nil {
		http.Error(w, "creating tiny url, "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := createPage{OriginalURL: inputUrl, TinyUrl: tinyURL}
	err = h.tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "executing template, "+err.Error(), http.StatusInternalServerError)
		return
	}
}
