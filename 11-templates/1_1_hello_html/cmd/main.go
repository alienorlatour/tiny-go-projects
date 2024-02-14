package main

import (
	"fmt"
	"net/http"
	"os"

	"learngo-pockets/templates/internal/handlers"
	"learngo-pockets/templates/internal/hlog"
)

const port = 8083

func main() {
	hlog.Set(os.Stdout)

	srv := handlers.New()

	addr := fmt.Sprintf(":%d", port)
	// log.Logger().Print("Listening on ", port, "...") FIXME log
	fmt.Println("Listening on ", port, "...") // FIXME log

	err := http.ListenAndServe(addr, srv.Router())
	if err != nil {
		panic(err)
	}
}
