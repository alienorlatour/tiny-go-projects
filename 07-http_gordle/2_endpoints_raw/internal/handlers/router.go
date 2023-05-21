package handlers

import (
	"net/http"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/handlers/newgame"
)

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(api.NewGameRoute, newgame.Handle)
	return mux
}
