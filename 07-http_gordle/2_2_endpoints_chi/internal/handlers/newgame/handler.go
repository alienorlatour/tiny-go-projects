package newgame

import (
	"encoding/json"
	"net/http"

	"learngo-pockets/httpgordle/api"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	apiGame := api.GameResponse{}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
