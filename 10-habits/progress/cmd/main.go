package main

import (
	"context"
	"fmt"
	"learngo-pockets/habits/internal/repository"
	"log/slog"
	"os"

	"learngo-pockets/habits/internal/server"
)

func main() {
	ctx := context.Background()
	srv := server.New(repository.New())

	// TODO: Catch Ctrl-C + defer graceful shutdown

	err := srv.Listen(ctx, 38804)
	if err != nil {
		slog.Error(fmt.Sprintf("Error while running the server: %s", err.Error()))
		os.Exit(1)
	}
}
