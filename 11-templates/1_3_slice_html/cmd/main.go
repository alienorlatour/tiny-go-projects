package main

import (
	"fmt"
	"log"
	"net/http"

	"learngo-pockets/templates/internal/handlers"
)

const port = 8083

func main() {
	srv := handlers.New()

	addr := fmt.Sprintf(":%d", port)
	log.Print("Listening on ", port, "...")

	err := http.ListenAndServe(addr, srv.Router())
	if err != nil {
		panic(err)
	}
}
