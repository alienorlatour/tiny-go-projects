package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"learngo-pockets/httpgordle/api"
)

func guessHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, ok := params["id"]
	if !ok {
		http.Error(writer, "missing the id of the game", http.StatusNotFound)
	}

	r := api.GuessRequest{}
	err := json.NewDecoder(request.Body).Decode(&r)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Store the guess for the game id

	log.Printf("guess %s for game id: %s", r.Guess, id)

	writer.WriteHeader(http.StatusNoContent)
}
