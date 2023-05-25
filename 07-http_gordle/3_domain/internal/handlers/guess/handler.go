package guess

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/handlers/apiconversion"
	"learngo-pockets/httpgordle/internal/session"
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

	game := guess(id, r)

	apiGame := apiconversion.ToAPIResponse(game)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}

func guess(id string, r api.GuessRequest) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}
