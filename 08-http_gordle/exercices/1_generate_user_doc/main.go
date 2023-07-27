package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"learngo-pockets/httpgordle/internal/handlers"
)

func main() {
	addr := ":8080"

	log.Print("Listening on ", addr, "...")

	r := handlers.NewRouter()

	// For each handler, print its method and its route.
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	// Start the server.
	err := http.ListenAndServe(addr, r)
	if err != nil {
		panic(err)
	}
}
