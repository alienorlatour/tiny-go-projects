package handlers

import (
	"github.com/go-chi/chi/v5"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/handlers/getstatus"
	"learngo-pockets/httpgordle/internal/handlers/guess"
	"learngo-pockets/httpgordle/internal/handlers/newgame"
	"learngo-pockets/httpgordle/internal/repository"
)

// NewRouter returns a router that listens for requests to the following endpoints:
//   - Create a new game;
//   - Get the status of a game;
//   - Make a guess in a game.
//
// The provided router is ready to serve.
func NewRouter(gr *repository.GameRepository) chi.Router {
	r := chi.NewRouter()

	r.Post(api.NewGamePath, newgame.Handler(gr))    // curl -X POST -v http://localhost:9090/games
	r.Get(api.GetStatusPath, getstatus.Handler(gr)) // curl -X GET -v http://localhost:9090/games/1682279480
	r.Put(api.GuessPath, guess.Handler(gr))         // curl -X PUT -v http://localhost:9090/games/1682279480 -d '{"value":"faune"}'

	return r
}
