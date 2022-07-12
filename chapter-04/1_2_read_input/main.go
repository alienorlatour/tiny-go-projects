package main

import (
	"fmt"

	"github.com/ablqk/tiny-go-projects/chapter-04/1_1_basic_main/gordle"
)

func main() {
	fmt.Println("Welcome to Gordle!")

	g := gordle.New()
	g.Play()
}
