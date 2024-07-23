package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed assets/*
//go:embed index.html
//go:embed wasm_exec.js
var assets embed.FS

func main() {
	fs := http.FileServer(http.FS(assets))
	http.Handle("/", fs)

	log.Print("Listening on :30001...")
	err := http.ListenAndServe(":30001", nil)

	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
