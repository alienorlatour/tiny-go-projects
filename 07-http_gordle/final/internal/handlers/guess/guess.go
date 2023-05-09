package guess

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/domain"
	"learngo-pockets/httpgordle/internal/gordle"
	"learngo-pockets/httpgordle/internal/handlers/apiconversion"
	"learngo-pockets/httpgordle/internal/repository"
)

type gameGuesser interface {
	Find(domain.GameID) (domain.Game, error)
	Update(domain.GameID, domain.Game) error
}

// Handler returns a handler for guess requests.
func Handler(repo gameGuesser) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Read the Game ID from the query parameters.
		id := chi.URLParam(request, api.GameID)
		if id == "" {
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
			case errors.Is(err, repository.ErrNotFound):
				http.Error(writer, err.Error(), http.StatusNotFound)
			case errors.Is(err, gordle.ErrInvalidGuess):
				http.Error(writer, err.Error(), http.StatusBadRequest)
			case errors.Is(err, domain.ErrGameOver):
				http.Error(writer, err.Error(), http.StatusForbidden)
			default:
				http.Error(writer, err.Error(), http.StatusInternalServerError)
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
	if game.AttemptsLeft == 0 || game.Status == domain.StatusWon {
		return domain.Game{}, domain.ErrGameOver
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
		Feedback: feedback.String(),
	})

	game.AttemptsLeft -= 1

	switch {
	case feedback.GameWon():
		game.Status = domain.StatusWon
	case game.AttemptsLeft == 0:
		game.Status = domain.StatusLost
	default:
		// Should be already set.
		game.Status = domain.StatusPlaying
	}

	// Update game status
	err = repo.Update(id, game)
	if err != nil {
		return domain.Game{}, fmt.Errorf("unable to save play: %w", err)
	}

	return game, nil
}
