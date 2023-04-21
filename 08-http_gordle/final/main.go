package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/game", newGameHandler)          // curl -X POST -v http://localhost:9090/game
	r.HandleFunc("/game/{id}", getStatus)          // curl -X GET -v http://localhost:9090/game/123
	r.HandleFunc("/game/{id}/guess", guessHandler) // curl -X POST -v http://localhost:9090/game/123/guess -d '{"value":"faune"}'
	log.Fatal(http.ListenAndServe(":9090", r))
}
