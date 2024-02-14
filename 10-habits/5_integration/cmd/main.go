package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"learngo-pockets/habits/internal/repository"
	"learngo-pockets/habits/internal/server"
	hlog "learngo-pockets/habits/log"
)

const port = 28710

func main() {
	// SIGINT is sent to the process when Ctrl-C is pressed while its running.
	// SIGTERM is a signal tools such as Kubernetes send to a container to shut it down.
	// SIGKILL is a signal sent to kill a process. It can't be caught.
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Set the writing output of our logger.
	hlog.Set(os.Stdout)

	srv := server.New(os.Stdout, repository.New())

	err := srv.Listen(ctx, port)
	if err != nil {
		log.Fatalf("Error while running the server: %s", err.Error())
	}
}
