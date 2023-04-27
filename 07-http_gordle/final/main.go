package main

import (
	"fmt"
	"log"
	"net/http"

	"learngo-pockets/httpgordle/internal/handlers"
	"learngo-pockets/httpgordle/internal/repository"
)

const port = 9090

func main() {
	gr := repository.New()

	r := handlers.NewRouter(gr)
	log.Printf("starting router on localhost:%d...", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	log.Fatal(err)
}
