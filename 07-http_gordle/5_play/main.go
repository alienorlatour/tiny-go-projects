package main

import (
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers"
	"learngo-pockets/httpgordle/internal/repository"
)

func main() {
	db := repository.New()

	err := http.ListenAndServe(":8080", handlers.NewRouter(db))
	if err != nil {
		panic(err)
	}
}
