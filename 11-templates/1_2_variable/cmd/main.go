package main

import (
	"fmt"
	"net/http"

	"learngo-pockets/templates/internal/handlers"
)

const port = 8083

func main() {
	// hlog.Set(os.Stdout) FIXME log

	srv := handlers.New()

	addr := fmt.Sprintf(":%d", port)
	// log.Logger().Print("Listening on ", port, "...") FIXME log
	fmt.Println("Listening on ", port, "...") // FIXME log

	err := http.ListenAndServe(addr, srv.Router())
	if err != nil {
		panic(err)
	}
}
