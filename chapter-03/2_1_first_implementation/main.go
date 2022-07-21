package main

import (
	"tiny-go-projects/chapter03/2_1_first_implementation/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// This produces nothing
	l.Debug("hello")
}
