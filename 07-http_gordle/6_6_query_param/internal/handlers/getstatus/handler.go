package getstatus

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/repository"
	"learngo-pockets/httpgordle/internal/session"
)

// gameFinder finds a game in the storage.
type gameFinder interface {
	Find(id session.GameID) (session.Game, error)
}

// Handler returns the handler for the game creation endpoint.
func Handler(finder gameFinder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, api.GameID)
		if id == "" {
			http.Error(w, "missing the id of the game", http.StatusNotFound)
			return
		}

		game, err := finder.Find(session.GameID(id))
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				http.Error(w, "this game does not exist yet", http.StatusNotFound)
				return
			}

			log.Printf("cannot fetch game %s: %s", id, err)
			http.Error(w, "failed to fetch game", http.StatusInternalServerError)
			return
		}

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			http.Error(w, "failed to write response", http.StatusInternalServerError)
			return
		}
	}
}
