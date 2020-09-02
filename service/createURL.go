package service

import (
	"fmt"
	"github.com/Matemateg/tinyurl/database"
	"math/rand"
	"net/url"
	"time"
)

type createUrl struct {
	dBase *database.DB
}

func NewCreateUrl(dBase *database.DB) *createUrl {
	return &createUrl{dBase: dBase}
}

func (c *createUrl) CreateURL(host string, inputUrl string) (string, error) {
	originalURL, err := url.Parse(inputUrl)
	if err != nil {
		return "", fmt.Errorf("parsing url, %w", err)
	}
	if originalURL.Scheme == "" {
		originalURL.Scheme = "http"
	}

	const tinyUrlLen = 7
	tinyUrlPath := fmt.Sprintf("/t/%v", generateRandomString(tinyUrlLen))

	err = c.dBase.Set(originalURL.String(), tinyUrlPath)
	if err != nil {
		return "", fmt.Errorf("setting to db, %w", err)
	}

	tinyURL := url.URL{}
	tinyURL.Scheme = "http"
	tinyURL.Host = host
	tinyURL.Path = tinyUrlPath

	return tinyURL.String(), nil
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
