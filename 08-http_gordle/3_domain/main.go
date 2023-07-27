package main

import (
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers"
)

func main() {
	addr := ":8080"

	log.Print("Listening on ", addr, "...")

	// Start the server.
	err := http.ListenAndServe(addr, handlers.NewRouter())
	if err != nil {
		panic(err)
	}
}
