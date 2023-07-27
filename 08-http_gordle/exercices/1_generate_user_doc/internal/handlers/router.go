package handlers

import (
	"github.com/go-chi/chi/v5"
	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/handlers/newgame"
)

// NewRouter returns a router that listens for requests to the following endpoints:
//   - Create a new game;
//
// The provided router is ready to serve.
func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Post(api.NewGameRoute, newgame.Handle)

	return r
}
