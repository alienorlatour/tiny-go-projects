package newgame

import (
	"encoding/json"
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers/apiconversion"
	"learngo-pockets/httpgordle/internal/session"
)

type gameAdder interface {
	Add(game session.Game) error
}

// Handler returns the handler for the game creation endpoint.
func Handler(repo gameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		game := createGame()

		apiGame := apiconversion.ToAPIResponse(game)

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			http.Error(w, "failed to write response", http.StatusInternalServerError)
		}
	}
}

func createGame() session.Game {
	return session.Game{}
}
