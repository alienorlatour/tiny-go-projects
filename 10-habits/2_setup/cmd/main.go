package main

import (
	"log"
	"os"

	"learngo-pockets/habits/internal/server"
	hlog "learngo-pockets/habits/log"
)

const port = 28710

func main() {
	// Set the writing output of our logger.
	hlog.Set(os.Stdout)

	srv := server.New()

	err := srv.Listen(port)
	if err != nil {
		log.Fatalf("Error while running the server: %s", err.Error())
	}
}
