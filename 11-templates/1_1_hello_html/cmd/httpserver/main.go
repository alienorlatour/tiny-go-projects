package main

import (
	"fmt"
	"net/http"
	"os"

	"learngo-pockets/templates/internal/handlers"
	"learngo-pockets/templates/internal/log"
)

const port = 8083

func main() {
	lgr := log.New(os.Stdout)

	srv := handlers.New(lgr)

	addr := fmt.Sprintf(":%d", port)
	lgr.Logf("Listening on %d...", port)

	err := http.ListenAndServe(addr, srv.Router())
	if err != nil {
		panic(err)
	}
}
