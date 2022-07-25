package main

import (
	"fmt"

	"github.com/ablqk/tiny-go-projects/chapter-04/1_3_separate_read_check/02_reader_gordle_param/gordle"
)

func main() {
	fmt.Println("Welcome to Gordle!")

	g := gordle.New()
	fmt.Println(g.Play())
}
