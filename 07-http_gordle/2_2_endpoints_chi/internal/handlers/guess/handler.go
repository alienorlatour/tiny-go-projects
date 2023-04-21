package guess

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"learngo-pockets/httpgordle/api"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, api.GameID)
	if id == "" {
		http.Error(w, "missing the id of the game", http.StatusNotFound)
		return
	}

	// Read the request, containing the guess, from the body of the input.
	r := api.GuessRequest{}
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	apiGame := api.GameResponse{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
