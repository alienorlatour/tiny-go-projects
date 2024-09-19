package guess

import (
	"encoding/json"
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/session"
)

type gameGuesser interface {
	Find(session.GameID) (session.Game, error)
	Update(game session.Game) error
}

// Handler returns the handler for the guess endpoint.
func Handler(guesser gameGuesser) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
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

		game := guess(id, r, guesser)

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}
	}
}

func guess(id string, r api.GuessRequest, db gameGuesser) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}
