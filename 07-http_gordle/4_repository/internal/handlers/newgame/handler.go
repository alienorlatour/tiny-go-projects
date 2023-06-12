package newgame

import (
	"encoding/json"
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers/apiconversion"
	"learngo-pockets/httpgordle/internal/session"
)

type gameAdder interface {
	Add(game session.Game) error
}

// Handler returns the handler for the game creation endpoint.
func Handler(db gameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		game := createGame(db)

		apiGame := apiconversion.ToAPIResponse(game)

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}
	}
}

func createGame(db gameAdder) session.Game {
	return session.Game{}
}
