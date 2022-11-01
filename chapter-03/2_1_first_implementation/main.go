package main

import (
	"github.com/ablqk/tiny-go-projects/chapter-03/2_1_first_implementation/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// This produces nothing
	l.Debug("hello")
}
