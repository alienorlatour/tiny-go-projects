package handlers

import (
	"github.com/go-chi/chi/v5"

	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/handlers/getstatus"
	"learngo-pockets/httpgordle/internal/handlers/newgame"
)

// NewRouter returns a router that listens for requests to the following endpoints:
//   - Create a new game;
//   - Get the status of a game;
//   - Make a guess in a game.
//
// The provided router is ready to serve.
func NewRouter() chi.Router {
	r := chi.NewRouter()

	// Register each endpoint.
	r.Post(api.NewGameRoute, newgame.Handle)
	r.Get(api.GetStatusRoute, getstatus.Handle)

	return r
}
