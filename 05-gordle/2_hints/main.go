package main

import (
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-04/2_hints/gordle"
)

const maxAttempts = 6

func main() {
	solution := "hello"

	g := gordle.New(os.Stdin, solution, maxAttempts)

	g.Play()
}
