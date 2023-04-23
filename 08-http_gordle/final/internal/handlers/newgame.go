package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"learngo-pockets/httpgordle/api"
)

// newGameHandler is the HTTP implementation of the endpoint that creates a new game.
func newGameHandler(writer http.ResponseWriter, _ *http.Request) {
	// Create the identified for a game.
	// rand isn't good enough for creating IDs, but it'll do for now.
	id := rand.Int()

	response := api.GameResponse{
		ID:           fmt.Sprint(id),
		AttemptsLeft: 0, // TODO: create a constant
		Guesses:      []api.Guess{},
	}

	// Header should be set before the writer.Write call.
	writer.WriteHeader(http.StatusCreated)

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, "failed to write response", http.StatusInternalServerError)
	}
}
