package newgame

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"

	"learngo-pockets/httpgordle/internal/gordle"
	"learngo-pockets/httpgordle/internal/handlers/apiconversion"
	"learngo-pockets/httpgordle/internal/session"
)

type gameCreator interface {
	Add(session.Game) error
}

// Handler returns the handler for the game creation endpoint.
func Handler(repo gameCreator) http.HandlerFunc {
	return func(writer http.ResponseWriter, _ *http.Request) {
		game, err := create(repo)
		if err != nil {
			log.Printf("unable to create a new game: %s", err)
			http.Error(writer, "failed to create a new game", http.StatusInternalServerError)
			return
		}

		apiGame := apiconversion.ToAPIResponse(game)

		// Header should be set before the writer.Write call.
		writer.WriteHeader(http.StatusCreated)

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(apiGame)
		if err != nil {
			http.Error(writer, "failed to write response", http.StatusInternalServerError)
		}
	}
}

const maxAttempts = 5

func create(repo gameCreator) (session.Game, error) {
	corpus, err := gordle.ReadCorpus("corpus/english.txt")
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to read corpus: %w", err)
	}

	if len(corpus) == 0 {
		return session.Game{}, gordle.ErrEmptyCorpus
	}

	word, err := gordle.PickRandomWord(corpus)
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to pick random word: %w", err)
	}

	game, err := gordle.New(word)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to create a new gordle game")
	}

	idInt, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt))
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to generate a random id")
	}

	id := session.GameID(fmt.Sprintf("%d", idInt))
	g := session.Game{
		ID:           id,
		Gordle:       *game,
		AttemptsLeft: maxAttempts,
		Guesses:      []session.Guess{},
		Status:       session.StatusPlaying,
	}

	err = repo.Add(g)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to save the new game")
	}

	return g, nil
}
