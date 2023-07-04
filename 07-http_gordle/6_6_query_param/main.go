package main

import (
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers"
	"learngo-pockets/httpgordle/internal/repository"
)

func main() {
	db := repository.New()

	addr := ":9090"

	log.Print("Listening on ", addr, "...")

	err := http.ListenAndServe(addr, handlers.NewRouter(db))
	if err != nil {
		panic(err)
	}
}
