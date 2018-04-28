package handlers

import (
	"encoding/json"
	"net/http"
)

type searchRequest struct {
	Query string `json:"query"`
}

// Search is an http handler for our microservice
type Search struct {
	Data string `json:"data"`
}

func (s *Search) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := searchRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil || len(request.Query) < 1 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

}
