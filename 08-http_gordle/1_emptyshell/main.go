package main

import (
	"log"
	"net/http"
)

func main() {
	addr := ":8080"

	log.Print("Listening on ", addr, "...")

	// Start the server.
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
