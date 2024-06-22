package getstatus

import (
	"encoding/json"
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/api"
)

// Handle is the handler for the status retrieval endpoint.
func Handle(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue(api.GameID)
	if id == "" {
		http.Error(w, "missing the id of the game", http.StatusBadRequest)
		return
	}

	apiGame := api.GameResponse{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		// The header has already been set. Nothing much we can do here.
		log.Printf("failed to write response: %s", err)
	}
}
