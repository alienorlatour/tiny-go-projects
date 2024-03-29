package handlers

import (
	"github.com/go-chi/chi/v5"

	"learngo-pockets/httpgordle/internal/api"
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
func NewRouter(db *repository.GameRepository) chi.Router {
	r := chi.NewRouter()

	r.Post(api.NewGameRoute, newgame.Handle(db))     // curl -X POST -v http://localhost:9090/games?lang="en"
	r.Get(api.GetStatusRoute, getstatus.Handler(db)) // curl -X GET -v http://localhost:9090/games/1682279480
	r.Put(api.GuessRoute, guess.Handle(db))          // curl -X PUT -v http://localhost:9090/games/1682279480 -d '{"guess":"faune"}'

	return r
}
