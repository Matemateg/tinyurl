package handlers

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Matemateg/tinyurl/database"
)

type redirectUrl struct{
	dBase *database.DB
}

func NewRedirectUrl(dBase *database.DB) *redirectUrl {
	return &redirectUrl{dBase: dBase}
}

func (rd *redirectUrl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tinyUrl := r.URL.Path
	originalUrl, err := rd.dBase.GetOriginal(tinyUrl)
	if err == database.ErrNotFound {
		fmt.Fprint(w, "Url не существует")
		return
	}
	if err != nil {
		fmt.Fprint(w, "Что-то пошло не так")
		log.Print(err)
		return
	}
	http.Redirect(w, r, originalUrl, 302)
}