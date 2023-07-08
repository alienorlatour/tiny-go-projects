package newgame

import (
	"encoding/json"
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/session"
)

type gameAdder interface {
	Add(game session.Game) error
}

// Handler returns the handler for the game creation endpoint.
func Handler(adder gameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		game := createGame(adder)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		apiGame := api.ToGameResponse(game)
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
