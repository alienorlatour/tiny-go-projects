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
	lgr := log.New(os.Stdout, log.Info)

	db := repository.New(lgr)

	srv := server.New(db, lgr)

	err := srv.ListenAndServe(port)
	if err != nil {
		lgr.Logf(log.Error, "Error while running the server: %s", err.Error())
	}
}
