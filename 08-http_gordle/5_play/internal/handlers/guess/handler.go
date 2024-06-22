package guess

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/gordle"
	"learngo-pockets/httpgordle/internal/repository"
	"learngo-pockets/httpgordle/internal/session"
)

type gameGuesser interface {
	Find(session.GameID) (session.Game, error)
	Update(game session.Game) error
}

// Handler returns the handler for the guess endpoint.
func Handler(guesser gameGuesser) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Read the Game ID from the path parameters.
		id := req.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing the id of the game", http.StatusBadRequest)
			return
		}

		// Read the request, containing the guess, from the body of the input.
		r := api.GuessRequest{}
		err := json.NewDecoder(req.Body).Decode(&r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		game, err := guess(session.GameID(id), r.Guess, guesser)
		if err != nil {
			switch {
			case errors.Is(err, repository.ErrNotFound):
				http.Error(w, err.Error(), http.StatusNotFound)
			case errors.Is(err, gordle.ErrInvalidGuessLength):
				http.Error(w, err.Error(), http.StatusBadRequest)
			case errors.Is(err, session.ErrGameOver):
				http.Error(w, err.Error(), http.StatusForbidden)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}
	}
}

func guess(id session.GameID, guess string, db gameGuesser) (session.Game, error) {
	// Does the game exist?
	game, err := db.Find(id)
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to find game: %w", err)
	}

	// Are plays still allowed?
	if game.AttemptsLeft == 0 || game.Status == session.StatusWon {
		return session.Game{}, session.ErrGameOver
	}

	// What does Gordle say about this guess ?
	feedback, err := game.Gordle.Play(guess)
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to play move: %w", err)
	}

	log.Printf("guess %v is valid in game %s", guess, id)

	// Record the play.
	game.Guesses = append(game.Guesses, session.Guess{
		Word:     guess,
		Feedback: feedback.String(),
	})

	game.AttemptsLeft -= 1

	switch {
	case feedback.GameWon():
		game.Status = session.StatusWon
	case game.AttemptsLeft == 0:
		game.Status = session.StatusLost
	default:
		// Should be already set.
		game.Status = session.StatusPlaying
	}

	// Update game
	err = db.Update(game)
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to save play: %w", err)
	}

	return game, nil
}
