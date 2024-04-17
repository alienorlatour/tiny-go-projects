package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/", fs)

	log.Print("Listening on :30001...")
	err := http.ListenAndServe(":30001", nil)

	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
