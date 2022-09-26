package main

import (
	"bufio"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-04/1_3_isolate_validate/gordle"
)

func main() {
	g := gordle.New(bufio.NewReader(os.Stdin))
	g.Play()
}
