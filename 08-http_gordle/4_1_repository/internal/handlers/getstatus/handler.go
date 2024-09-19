package getstatus

import (
	"encoding/json"
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/session"
)

// Handler returns the handler for the status retrieval endpoint.
// The repo parameter will be more clearly defined in the next section.
func Handler(repo interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing the id of the game", http.StatusBadRequest)
			return
		}

		game := getGame(id)

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}
	}
}

func getGame(id string) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}
