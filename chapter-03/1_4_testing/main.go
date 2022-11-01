package main

import (
	"github.com/ablqk/tiny-go-projects/chapter-03/1_4_testing/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// This produces nothing
	l.Debugf("hello")
}
