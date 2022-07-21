package main

import (
	"tiny-go-projects/chapter03/1_3_new/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// This produces nothing
	l.Debug("hello")
}
