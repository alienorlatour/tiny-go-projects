package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-04/2_feedback/gordle"
)

func main() {
	fmt.Println("Welcome to Gordle!")

	solution := "hello"
	g := gordle.New(bufio.NewReader(os.Stdin), []rune(solution), 6)
	g.Play()
}
