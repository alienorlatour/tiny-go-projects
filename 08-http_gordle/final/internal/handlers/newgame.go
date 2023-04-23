package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"learngo-pockets/httpgordle/api"
)

func newGameHandler(writer http.ResponseWriter, _ *http.Request) {
	id := time.Now().Unix()

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
