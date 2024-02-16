package main

import (
	"os"

	"learngo-pockets/habits/internal/log"
	"learngo-pockets/habits/internal/repository"
	"learngo-pockets/habits/internal/server"
)

const port = 28710

func main() {
	// Set the writing output of our logger.
	log.Set(os.Stdout)

	db := repository.New()

	srv := server.New(db)

	err := srv.Listen(port)
	if err != nil {
		log.Fatalf("Error while running the server: %s", err.Error())
	}
}
