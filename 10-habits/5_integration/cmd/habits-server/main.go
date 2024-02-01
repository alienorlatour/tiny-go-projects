package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"learngo-pockets/habits/internal/log"
	"learngo-pockets/habits/internal/repository"
	"learngo-pockets/habits/internal/server"
)

const port = 28710

func main() {
	// SIGINT is sent to the process when Ctrl-C is pressed while its running.
	// SIGTERM is a signal tools such as Kubernetes send to a container to shut it down.
	// SIGKILL is a signal sent to kill a process. It can't be caught.
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Set the writing output of our logger.
	lgr := log.New(os.Stdout)

	db := repository.New(lgr)
	srv := server.New(db, lgr)

	err := srv.ListenAndServe(ctx, port)
	if err != nil {
		lgr.Logf("Error while running the server: %s", err.Error())
		os.Exit(1)
	}
}
