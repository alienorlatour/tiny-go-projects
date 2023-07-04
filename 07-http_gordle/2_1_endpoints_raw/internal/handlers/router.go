package handlers

import (
	"net/http"

	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/handlers/newgame"
)

// Mux implements the http router to the endpoints.
func Mux() *http.ServeMux {
	mux := http.NewServeMux()

	// Register the newGame endpoint.
	mux.HandleFunc(api.NewGameRoute, newgame.Handle)

	return mux
}
