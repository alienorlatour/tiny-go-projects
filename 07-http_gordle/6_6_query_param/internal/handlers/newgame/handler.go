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

// Handle returns the handler for the game creation endpoint.
func Handle(adder gameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get(api.Lang)
		if len(lang) > 0 {
			// TODO create a game in the chosen language
			fmt.Println(lang)
		}

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
			http.Error(w, "failed to write response", http.StatusInternalServerError)
		}
	}
}

const maxAttempts = 5

var corpusPath = "corpus/english.txt"

func createGame(db gameAdder) (session.Game, error) {
	corpus, err := gordle.ReadCorpus(corpusPath)
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to read corpus: %w", err)
	}

	game, err := gordle.New(corpus)
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
