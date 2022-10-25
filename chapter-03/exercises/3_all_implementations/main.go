package main

import (
	"tiny-go-projects/chapter03/exercises/3_all_implementations/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// This produces nothing
	l.Debugf("hello")
}
