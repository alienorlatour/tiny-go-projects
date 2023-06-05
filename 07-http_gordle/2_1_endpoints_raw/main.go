package main

import (
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers"
)

func main() {
	addr := ":8080"

	log.Print("Listening on ", addr, "...")

	err := http.ListenAndServe(addr, handlers.Mux())
	if err != nil {
		panic(err)
	}
}
