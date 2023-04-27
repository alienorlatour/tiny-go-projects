package getstatus

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/domain"
	"learngo-pockets/httpgordle/internal/handlers"
	"learngo-pockets/httpgordle/internal/repository"
)

type gameFinder interface {
	Find(id domain.GameID) (domain.Game, error)
}

// Handler returns the handler for the game finder endpoint.
func Handler(repo gameFinder) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params[api.GameID]
		log.Printf("retrieve status from id: %v", id)

		game, err := repo.Find(domain.GameID(id))
		if err != nil {
			switch {
			case errors.Is(err, repository.ErrNotFound):
				writer.WriteHeader(http.StatusNotFound)
			default:
				writer.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		// TODO: retrieve status from game id
		apiGame := handlers.ToAPI(game)

		writer.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(writer).Encode(apiGame)
		if err != nil {
			http.Error(writer, "failed to write response", http.StatusInternalServerError)
		}
	}
}
