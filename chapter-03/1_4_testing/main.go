package main

import (
	"tiny-go-projects/chapter03/1_4_testing/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// This produces nothing
	l.Debug("hello")
}
