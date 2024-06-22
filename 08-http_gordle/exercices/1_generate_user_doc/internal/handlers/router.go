package handlers

import (
	"net/http"

	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/handlers/newgame"
)

// NewRouter returns a router that listens for requests to the following endpoints:
//   - Create a new game;
//
// The provided router is ready to serve.
func NewRouter() *http.ServeMux {
	r := http.NewServeMux()

	// Register each endpoint.
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handle)

	return r
}
