package newgame

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/domain"
	"learngo-pockets/httpgordle/internal/gordle"
)

type gameCreator interface {
	Add(game domain.Game) error
}

// Handler returns the handler for the game creation endpoint.
func Handler(repo gameCreator) http.HandlerFunc {
	return func(writer http.ResponseWriter, _ *http.Request) {
		game, err := create(repo)
		if err != nil {
			log.Printf("unable to create a new game: %s", err)
			http.Error(writer, "failed to create a new game", http.StatusInternalServerError)
		}

		response := api.GameResponse{
			ID:           string(game.ID),
			AttemptsLeft: byte(game.Gordle.MaxAttempts), // TODO
			Guesses:      []api.Guess{},
		}

		// Header should be set before the writer.Write call.
		writer.WriteHeader(http.StatusCreated)

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(response)
		if err != nil {
			http.Error(writer, "failed to write response", http.StatusInternalServerError)
		}
	}
}

func create(repo gameCreator) (domain.Game, error) {
	game, err := gordle.New([]string{"LOGIN", "HELLO"})
	if err != nil {
		return domain.Game{}, fmt.Errorf("failed to create a new gordle game")
	}

	id := domain.GameID(fmt.Sprintf("%d", rand.Int()))
	g := domain.Game{ID: id, Gordle: *game}

	err = repo.Add(g)
	if err != nil {
		return domain.Game{}, fmt.Errorf("failed to save the new game")
	}

	return g, nil
}
