package main

import (
	"net/http"

	"learngo-pockets/httpgordle/internal"
)

func main() {
	err := http.ListenAndServe(":8080", internal.Mux())
	if err != nil {
		panic(err)
	}
}
