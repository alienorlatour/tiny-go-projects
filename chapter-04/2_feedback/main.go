package main

import (
	"bufio"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-04/2_feedback/gordle"
)

const maxAttempts = 6

func main() {
	solution := "hello"
	g := gordle.New(bufio.NewReader(os.Stdin), []rune(solution), maxAttempts)
	g.Play()
}
