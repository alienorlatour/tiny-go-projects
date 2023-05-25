package newgame

import (
	"encoding/json"
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers/apiconversion"
	"learngo-pockets/httpgordle/internal/session"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	game := createGame()

	apiGame := apiconversion.ToAPIResponse(game)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}

func createGame() session.Game {
	return session.Game{}
}
