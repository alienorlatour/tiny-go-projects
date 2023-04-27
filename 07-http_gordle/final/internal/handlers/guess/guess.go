package guess

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/domain"
	"learngo-pockets/httpgordle/internal/repository"
)

type gameGuesser interface {
	Guess(id domain.GameID, word string) (domain.Game, error)
}

func Handler(repo gameGuesser) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id, ok := params[api.GameID]
		if !ok {
			http.Error(writer, "missing the id of the game", http.StatusNotFound)
		}

		r := api.GuessRequest{}
		err := json.NewDecoder(request.Body).Decode(&r)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		game, err := repo.Guess(domain.GameID(id), r.Guess)
		if err != nil {
			switch {
			case errors.Is(err, repository.ErrNotFound):
				writer.WriteHeader(http.StatusNotFound)
			default:
				writer.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		log.Printf("feedback: %s", game.Guesses[len(game.Guesses)-1].Feedback)

		writer.WriteHeader(http.StatusNoContent)
	}
}
