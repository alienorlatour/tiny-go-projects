package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"learngo-pockets/habits/internal/repository"
	"learngo-pockets/habits/internal/server"
)

const port = 28710

func main() {
	/*	ctx, cancel := context.WithCancel(context.Background())

		// Create a channel to receive OS signals
		sigChan := make(chan os.Signal, 1)
		defer close(sigChan)
		ctx, cancel := signal.NotifyContext(sigChan, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigChan)
		go func() {
			sig := <-sigChan
			switch sig {
			case syscall.SIGINT:
				log.Println("Server was stopped with a SIGINT call")
			case syscall.SIGTERM:
				log.Println("Server was stopped with a SIGTERM call")
			}

			// Graceful shutdown
			cancel()
		}()
	*/

	// SIGINT is sent to the process when Ctrl-C is pressed while its running.
	// SIGTERM is a signal tools such as Kubernetes send to a container to shut it down.
	// SIGKILL is a signal sent to kill a process. It can't be caught.
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	srv := server.New(repository.New())

	err := srv.Listen(ctx, port)
	if err != nil {
		log.Fatalf("Error while running the server: %s", err.Error())
	}
}
