package main

import (
	"log"
	"net/http"
)

func main() {
	addr := ":8080"

	log.Print("Listening on ", addr, "...")

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
