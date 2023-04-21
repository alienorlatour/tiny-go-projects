package main

import (
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers"
)

func main() {
	err := http.ListenAndServe(":8080", handlers.Mux())
	if err != nil {
		panic(err)
	}
}
