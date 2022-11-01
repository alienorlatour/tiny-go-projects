package main

import (
	"github.com/ablqk/tiny-go-projects/chapter-03/1_2_object_oriented/pocketlog"
)

func main() {
	lgr := pocketlog.Logger{}

	// This produces nothing
	lgr.Debugf("Make the zero (%d) value useful.", 0)
}
