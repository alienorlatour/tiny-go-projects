package main

import (
	"fmt"
	"log/slog"
	"os"

	"learngo-pockets/habits/internal/repository"
	"learngo-pockets/habits/internal/server"
)

const port = 38804

func main() {
	srv := server.New(repository.New())

	// TODO: Catch Ctrl-C + defer graceful shutdown

	err := srv.Listen(port)
	if err != nil {
		slog.Error(fmt.Sprintf("Error while running the server: %s", err.Error()))
		os.Exit(1)
	}
}
