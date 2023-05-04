package handlers

import (
	"github.com/go-chi/chi/v5"

	"learngo-pockets/httpgordle/api"
)

// NewRouter returns a router that listens for requests to the following endpoints:
//   - Create a new game;
//   - Get the status of a game;
//   - Make a guess in a game.
//
// The provided router is ready to serve.
func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Post(api.NewGamePath, newGameHandler) // curl -X POST -v http://localhost:9090/games
	r.Get(api.GetStatusPath, getStatus)     // curl -X GET -v http://localhost:9090/games/1682279480
	r.Put(api.GuessPath, guessHandler)      // curl -X PUT -v http://localhost:9090/games/1682279480 -d '{"value":"faune"}'

	return r
}
