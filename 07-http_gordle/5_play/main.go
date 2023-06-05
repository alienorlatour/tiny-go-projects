package main

import (
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers"
	"learngo-pockets/httpgordle/internal/repository"
)

func main() {
	// Create the games' repository.
	db := repository.New()

	addr := ":8080"

	log.Print("Listening on ", addr, "...")

	// Start the server.
	err := http.ListenAndServe(addr, handlers.NewRouter(db))
	if err != nil {
		panic(err)
	}
}
