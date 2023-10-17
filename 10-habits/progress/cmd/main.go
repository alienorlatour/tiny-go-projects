package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"learngo-pockets/habits/internal/server"
)

func main() {
	ctx := context.Background()
	srv := server.New()

	// TODO: Catch Ctrl-C + defer graceful shutdown

	err := srv.Listen(ctx, 38804)
	if err != nil {
		slog.Error(fmt.Sprintf("Error while running the server: %s", err.Error()))
		os.Exit(1)
	}
}
