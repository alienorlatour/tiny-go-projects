package main

import (
	"log"

	"learngo-pockets/habits/internal/server"
)

const port = 28710

func main() {
	srv := server.New()

	err := srv.Listen(port)
	if err != nil {
		log.Fatalf("Error while running the server: %s", err.Error())
	}
}
