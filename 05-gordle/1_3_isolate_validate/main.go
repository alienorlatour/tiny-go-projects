package main

import (
	"os"

	"learngo-pockets/gordle/gordle"
)

func main() {
	g := gordle.New(os.Stdin)
	g.Play()
}
