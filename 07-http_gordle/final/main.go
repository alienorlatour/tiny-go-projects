package main

import (
	"fmt"
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers"
)

const port = 9090

func main() {
	r := handlers.NewRouter()
	log.Printf("starting router on localhost:%d...", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	log.Fatal(err)
}
