package getstatus

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/domain"
	"learngo-pockets/httpgordle/internal/handlers"
)

type gameFinder interface {
	Find(id domain.GameID) *domain.Game
}

// Handler returns the handler for the game finder endpoint.
func Handler(repo gameFinder) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id := params[api.GameID]
		log.Printf("retrieve status from id: %v", id)

		game := repo.Find(domain.GameID(id))
		if game == nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		// TODO: retrieve status from game id
		apiGame := handlers.ToAPI(*game)

		writer.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(writer).Encode(apiGame)
		if err != nil {
			http.Error(writer, "failed to write response", http.StatusInternalServerError)
		}
	}
}
