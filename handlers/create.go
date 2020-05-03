package handlers

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
	"github.com/Matemateg/tinyurl/database"
)

type createPage struct{
	OriginalURL string
	TinyUrl string
}

type createUrl struct{
	dBase *database.DB
}

func NewCreateUrl(dBase *database.DB) *createUrl {
	return &createUrl{dBase: dBase}
}

func (c *createUrl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	inputUrl := r.FormValue("inputUrl")
	originalURL, err := url.Parse(inputUrl)
	if err != nil {
		fmt.Fprintf(w, "что-то пошло не так")
		log.Print(err)
		return
	}
	if originalURL.Scheme == "" {
		originalURL.Scheme = "http"
	}

	const tinyUrlLen = 10
	tinyUrlPath := fmt.Sprintf("/t/%v", generateRandomString(tinyUrlLen))

	err = c.dBase.Set(originalURL.String(), tinyUrlPath)
	if err != nil {
		fmt.Fprintf(w, "что-то пошло не так")
		log.Print(err)
		return
	}

	t, _ := template.ParseFiles("templates/create.html")

	tinyURL := url.URL{}
	tinyURL.Scheme = "http"
	tinyURL.Host = r.Host
	tinyURL.Path = tinyUrlPath

	data := createPage{OriginalURL: originalURL.String(), TinyUrl: tinyURL.String()}
	t.Execute(w, data)
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	alphabet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	randSrc := rand.NewSource(time.Now().UnixNano())
	result := make([]rune, length)
	for i := range result {
		result[i] = alphabet[randSrc.Int63()%int64(len(alphabet))]
	}
	return string(result)
}