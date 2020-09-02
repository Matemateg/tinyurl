package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UrlCreator interface {
	CreateURL(host string, inputUrl string) (string, error)
}

type creatingURL struct {
	service UrlCreator
}

func NewCreatingURL(service UrlCreator) *creatingURL {
	return &creatingURL{service: service}
}

type request struct {
	URL string `json:"url"`
}

type response struct {
	URL string `json:"url"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func (h *creatingURL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err := fmt.Errorf("parsing json, %w", err)
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	tinyURL, err := h.service.CreateURL(r.Host, req.URL)
	if err != nil {
		err := fmt.Errorf("creating tiny, %w", err)
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	resp := response{URL: tinyURL}
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		err := fmt.Errorf("creating json, %w", err)
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	errRes := errorResponse{Error: err.Error()}
	_ = json.NewEncoder(w).Encode(errRes)
}
