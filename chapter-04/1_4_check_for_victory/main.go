package main

import (
	"fmt"

	"github.com/ablqk/tiny-go-projects/chapter-04/1_4_check_for_victory/gordle"
)

func main() {
	fmt.Println("Welcome to Gordle!")

	solution := "hello"
	g := gordle.New([]rune(solution))
	g.Play()
}
