package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
)

//go:embed index.html
//go:embed main.wasm
//go:embed wasm_exec.js
var assets embed.FS

func main() {
	fs := http.FileServer(http.FS(assets))
	http.Handle("/", fs)

	log.Print("Listening on 127.0.0.1:30001...")
	err := http.ListenAndServe("127.0.0.1:30001", nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start server: %s", err)
		return
	}
}
