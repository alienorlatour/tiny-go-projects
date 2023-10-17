package main

import (
	"context"
	"fmt"
	"os"

	"learngo-pockets/habits/internal/server"
)

func main() {
	ctx := context.Background()
	srv := server.New()

	// TODO: Catch Ctrl-C + defer graceful shutdown

	err := srv.Listen(ctx, 38804)
	if err != nil {
		fmt.Printf("Error while running the server: %s", err.Error())
		os.Exit(1)
	}
}
