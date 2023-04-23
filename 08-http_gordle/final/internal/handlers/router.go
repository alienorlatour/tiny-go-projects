package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/games", newGameHandler).Methods(http.MethodPost)   // curl -X POST -v http://localhost:9090/games
	r.HandleFunc("/games/{id}", getStatus).Methods(http.MethodGet)    // curl -X GET -v http://localhost:9090/games/123
	r.HandleFunc("/games/{id}", guessHandler).Methods(http.MethodPut) // curl -X PUT -v http://localhost:9090/games/123 -d '{"value":"faune"}'

	return r
}
