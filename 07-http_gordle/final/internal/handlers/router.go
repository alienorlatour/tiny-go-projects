package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/handlers/newgame"
	"learngo-pockets/httpgordle/internal/repository"
)

// NewRouter returns a router that listens for requests to the following endpoints:
//   - Create a new game;
//   - Get the status of a game;
//   - Make a guess in a game.
//
// The provided router is ready to serve.
func NewRouter(gr *repository.GameRepository) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc(api.NewGamePath, newgame.Handler(gr)).Methods(http.MethodPost) // curl -X POST -v http://localhost:9090/games
	r.HandleFunc(api.GetStatusPath, getStatus).Methods(http.MethodGet)          // curl -X GET -v http://localhost:9090/games/1682279480
	r.HandleFunc(api.GuessPath, guessHandler).Methods(http.MethodPut)           // curl -X PUT -v http://localhost:9090/games/1682279480 -d '{"value":"faune"}'

	return r
}
