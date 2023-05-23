package getstatus

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/domain"
	"learngo-pockets/httpgordle/internal/handlers/apiconversion"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, api.GameID)
	if id == "" {
		http.Error(w, "missing the id of the game", http.StatusNotFound)
		return
	}

	game := getGame(id)

	apiGame := apiconversion.ToAPIResponse(game)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}

func getGame(id string) domain.Game {
	return domain.Game{
		ID: domain.GameID(id),
	}
}
