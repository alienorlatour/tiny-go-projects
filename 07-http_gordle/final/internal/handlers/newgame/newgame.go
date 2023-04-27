package newgame

import (
	"encoding/json"
	"net/http"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/domain"
)

type gameCreator interface {
	Create() domain.Game
}

// Handler returns the handler for the game creation endpoint.
func Handler(repo gameCreator) http.HandlerFunc {
	return func(writer http.ResponseWriter, _ *http.Request) {
		g := repo.Create()
		response := api.GameResponse{
			ID:           string(g.ID),
			AttemptsLeft: g.AttemptsLeft,
			Guesses:      []api.Guess{},
		}

		// Header should be set before the writer.Write call.
		writer.WriteHeader(http.StatusCreated)

		writer.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(response)
		if err != nil {
			http.Error(writer, "failed to write response", http.StatusInternalServerError)
		}
	}
}
