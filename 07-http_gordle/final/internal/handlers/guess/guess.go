package guess

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/domain"
	"learngo-pockets/httpgordle/internal/gordle"
	"learngo-pockets/httpgordle/internal/handlers/apiconversion"

	"github.com/gorilla/mux"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/repository"
)

type gameGuesser interface {
	Find(id domain.GameID) (domain.Game, error)
	Update(domain.GameID, domain.Game) error
}

// Handler returns a handler for guess requests.
func Handler(repo gameGuesser) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Read the Game ID from the query parameters.
		params := mux.Vars(request)
		id, ok := params[api.GameID]
		if !ok {
			http.Error(writer, "missing the id of the game", http.StatusNotFound)
		}

		// Read the request, containing the guess, from the body of the input.
		r := api.GuessRequest{}
		err := json.NewDecoder(request.Body).Decode(&r)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		game, err := play(repo, domain.GameID(id), r.Guess)
		if err != nil {
			switch {
			// TODO: I don't know where to put all the different errors. I want somewhere where they make sense.
			case errors.Is(err, repository.ErrNotFound):
				http.Error(writer, err.Error(), http.StatusNotFound)
			case errors.Is(err, gordle.ErrInvalidGuessLength):
				http.Error(writer, err.Error(), http.StatusBadRequest)
			}
			return
		}

		apiGame := apiconversion.ToAPIResponse(game)

		writer.WriteHeader(http.StatusAccepted)

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(apiGame)
		if err != nil {
			http.Error(writer, "failed to write response", http.StatusInternalServerError)
			return
		}
	}
}

func play(repo gameGuesser, id domain.GameID, guess string) (domain.Game, error) {
	// Does the game exist?
	game, err := repo.Find(id)
	if err != nil {
		return domain.Game{}, fmt.Errorf("unable to find game repository: %w", err)
	}

	// Are plays still allowed?
	if game.AttemptsLeft == 0 {
		return domain.Game{}, fmt.Errorf("no more plays allowed")
	}

	// What does Gordle say about this guess ?
	feedback, err := game.Gordle.Play(guess)
	if err != nil {
		return domain.Game{}, fmt.Errorf("unable to play move: %w", err)
	}

	log.Printf("Guess %v is valid in game %s", guess, id)

	// Record the play.
	game.Guesses = append(game.Guesses, domain.Guess{
		Word:     guess,
		Feedback: feedback,
	})
	game.AttemptsLeft -= 1

	// Update game status
	err = repo.Update(id, game)
	if err != nil {
		return domain.Game{}, fmt.Errorf("unable to save play: %w", err)
	}

	return game, nil
}
