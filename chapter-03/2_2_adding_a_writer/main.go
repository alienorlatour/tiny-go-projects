package main

import (
	"tiny-go-projects/chapter03/2_2_adding_a_writer/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// This produces nothing
	l.Debugf("hello")
}
