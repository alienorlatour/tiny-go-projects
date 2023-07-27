package newgame

import (
	"net/http"

	"learngo-pockets/httpgordle/internal/api"
)

// Handle is the handler for the game creation endpoint.
func Handle(w http.ResponseWriter, req *http.Request) {
	if req.Method != api.NewGameMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Creating a new game"))
}
