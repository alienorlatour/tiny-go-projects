package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type guessRequest struct {
	Value string `json:"value"`
}

func guessHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(request)
	id, ok := params["id"]
	if !ok {
		http.Error(writer, "missing the id of the game", http.StatusNotFound)
	}

	guess := guessRequest{}
	err := json.NewDecoder(request.Body).Decode(&guess)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("guess %s for game id: %s", guess.Value, id)

	writer.WriteHeader(http.StatusCreated)
}
