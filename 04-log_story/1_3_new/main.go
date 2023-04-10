package main

import (
	"github.com/ablqk/tiny-go-projects/chapter-03/1_3_new/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelDebug)

	// This produces nothing
	lgr.Debugf("Make the zero (%d) value useful.", 0)
}
