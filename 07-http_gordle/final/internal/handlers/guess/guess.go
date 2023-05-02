package guess

import (
	"encoding/json"
	"errors"
	"learngo-pockets/httpgordle/internal/domain"
	"net/http"

	"github.com/gorilla/mux"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/gordle"
	"learngo-pockets/httpgordle/internal/repository"
)

type gameGuesser interface {
	Find(id domain.GameID) (gordle.Game, error)
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

		game.Play()

		// TODO
		// game.log.Printf("feedback: %s", game.Guesses[len(game.Guesses)-1].Feedback)

		writer.WriteHeader(http.StatusNoContent)
	}
}
