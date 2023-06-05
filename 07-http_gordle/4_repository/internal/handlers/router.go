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
func NewRouter(db *repository.GameRepository) chi.Router {
	r := chi.NewRouter()

	// Register each endpoint.
	r.Post(api.NewGameRoute, newgame.Handler(db))
	r.Get(api.GetStatusRoute, getstatus.Handler(db))
	r.Put(api.GuessRoute, guess.Handler(db))

	return r
}
