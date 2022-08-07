package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-04/1_3_isolate_validate/gordle"
)

func main() {
	fmt.Println("Welcome to Gordle!")

	g := gordle.New(bufio.NewReader(os.Stdin))
	g.Play()
}
