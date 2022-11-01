package main

import (
	"github.com/ablqk/tiny-go-projects/chapter-03/exercises/3_all_implementations/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// This produces nothing
	l.Debugf("hello")
}
