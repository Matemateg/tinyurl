package handlers

import (
	"github.com/Matemateg/tinyurl/database"
	"net/http"
)

type redirectUrl struct {
	dBase *database.DB
}

func NewRedirectUrl(dBase *database.DB) *redirectUrl {
	return &redirectUrl{dBase: dBase}
}

func (rd *redirectUrl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tinyUrl := r.URL.Path
	originalUrl, err := rd.dBase.GetOriginal(tinyUrl)
	if err == database.ErrNotFound {
		http.Error(w, "Url не существует", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, "getting url, "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, originalUrl, http.StatusFound)
}
