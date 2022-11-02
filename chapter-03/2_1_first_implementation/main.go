package main

import (
	"github.com/ablqk/tiny-go-projects/chapter-03/2_1_first_implementation/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo)

	// This produces nothing
	lgr.Debugf("Make the zero (%d) value useful.", 0)
}
