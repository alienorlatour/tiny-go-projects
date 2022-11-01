package main

import (
	"github.com/ablqk/tiny-go-projects/chapter-03/2_2_adding_a_writer/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// This produces nothing
	l.Debugf("hello")
}
