package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"learngo-pockets/templates/internal/handlers"
)

const port = 8083

func main() {
	ctx := context.Background()

	srv := handlers.New(ctx, "localhost:28710")

	addr := fmt.Sprintf(":%d", port)
	log.Print("Listening on ", port, "...")

	err := http.ListenAndServe(addr, srv.Router())
	if err != nil {
		panic(err)
	}
}
