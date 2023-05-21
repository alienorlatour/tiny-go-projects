package main

import "net/http"

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
