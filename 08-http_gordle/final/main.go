package main

import (
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers"
)

func main() {
	r := handlers.NewRouter()
	log.Println("starting router on localhost:9090...")

	err := http.ListenAndServe(":9090", r)
	log.Fatal(err)
}
