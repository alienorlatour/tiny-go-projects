package newgame

import (
	"encoding/json"
	"net/http"

	"learngo-pockets/httpgordle/internal/api"
)

// Handle is the handler for the game creation endpoint.
func Handle(w http.ResponseWriter, req *http.Request) {
	apiGame := api.GameResponse{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
