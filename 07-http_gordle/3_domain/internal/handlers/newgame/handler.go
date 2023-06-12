package newgame

import (
	"encoding/json"
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers/apiconversion"
	"learngo-pockets/httpgordle/internal/session"
)

// Handle is the handler for the game creation endpoint.
func Handle(w http.ResponseWriter, req *http.Request) {
	game := createGame()

	apiGame := apiconversion.ToAPIResponse(game)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		// The header has already been set. Nothing much we can do here.
		log.Printf("failed to write response: %s", err)
	}
}

func createGame() session.Game {
	return session.Game{}
}
