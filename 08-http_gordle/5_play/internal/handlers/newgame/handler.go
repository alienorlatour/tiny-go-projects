package newgame

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/oklog/ulid/v2"

	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/gordle"
	"learngo-pockets/httpgordle/internal/session"
)

type gameAdder interface {
	Add(game session.Game) error
}

// Handler returns the handler for the game creation endpoint.
func Handler(adder gameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		game, err := createGame(adder)
		if err != nil {
			log.Printf("unable to create a new game: %s", err)
			http.Error(w, "failed to create a new game", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		apiGame := api.ToGameResponse(game)
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}
	}
}

const maxAttempts = 5

func createGame(db gameAdder) (session.Game, error) {
	corpus, err := gordle.ParseCorpus()
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to read corpus: %w", err)
	}

	if len(corpus) == 0 {
		return session.Game{}, gordle.ErrEmptyCorpus
	}

	solution, err := gordle.PickRandomWord(corpus)
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to pick a random solution: %w", err)
	}

	game, err := gordle.New(solution)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to create a new gordle game")
	}

	g := session.Game{
		ID:           session.GameID(ulid.Make().String()),
		Gordle:       *game,
		AttemptsLeft: maxAttempts,
		Guesses:      []session.Guess{},
		Status:       session.StatusPlaying,
	}

	err = db.Add(g)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to save the new game")
	}

	return g, nil
}
