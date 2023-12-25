package main

import (
	"log"

	"learngo-pockets/habits/internal/repository"
	"learngo-pockets/habits/internal/server"
)

const port = 28710

func main() {
	db := repository.New()

	srv := server.New(db)

	err := srv.Listen(port)
	if err != nil {
		log.Fatalf("Error while running the server: %s", err.Error())
	}
}
