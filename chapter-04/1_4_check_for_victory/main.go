package main

import (
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-04/1_4_check_for_victory/gordle"
)

const maxAttempts = 6

func main() {
	solution := "hello"
	g := gordle.New(os.Stdin, []rune(solution), maxAttempts)
	g.Play()
}
