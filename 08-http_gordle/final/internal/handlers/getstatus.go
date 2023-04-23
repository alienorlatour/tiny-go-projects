package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type getStatusResponse struct {
	ID           string   `json:"id"`
	AttemptsLeft string   `json:"attemptsLeft"`
	Guesses      []string `json:"guesses"`
}

func getStatus(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(request)
	id := params["id"]
	log.Printf("retrieve status from id: %v", id)

	// TODO: retrieve status from game id
	response := getStatusResponse{
		ID:           "123",
		AttemptsLeft: "2",
		Guesses:      []string{"hello", "sauna", "files"},
	}

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, "failed to write response", http.StatusInternalServerError)
	}
}
