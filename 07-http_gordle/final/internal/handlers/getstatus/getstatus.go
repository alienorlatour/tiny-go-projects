package getstatus

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/handlers/apiconversion"
	"learngo-pockets/httpgordle/internal/repository"
	"learngo-pockets/httpgordle/internal/session"
)

type gameFinder interface {
	Find(session.GameID) (session.Game, error)
}

// Handler returns the handler for the game finder endpoint.
func Handler(repo gameFinder) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, api.GameID)
		if id == "" {
			http.Error(writer, "missing the id of the game", http.StatusNotFound)
			return
		}
		log.Printf("retrieve status from id: %v", id)

		game, err := repo.Find(session.GameID(id))
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				writer.WriteHeader(http.StatusNotFound)
			} else {
				writer.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		apiGame := apiconversion.ToAPIResponse(game)

		writer.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(writer).Encode(apiGame)
		if err != nil {
			http.Error(writer, "failed to write response", http.StatusInternalServerError)
		}
	}
}
