package handlers

import (
	"net/http"

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
func NewRouter(db *repository.GameRepository) *http.ServeMux {
	r := http.NewServeMux()

	// Register each endpoint.
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handler(db))    // curl -X POST -v http://localhost:9090/games
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handler(db)) // curl -X GET -v http://localhost:9090/games/1682279480
	r.HandleFunc(http.MethodPut+" "+api.GuessRoute, guess.Handler(db))         // curl -X PUT -v http://localhost:9090/games/1682279480 -d '{"guess":"faune"}'

	return r
}
