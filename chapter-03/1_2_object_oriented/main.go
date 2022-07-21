package main

import (
	"tiny-go-projects/chapter03/1_2_object_oriented/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// This produces nothing
	l.Debug("hello")
}
