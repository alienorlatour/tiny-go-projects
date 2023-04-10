package main

import (
	"github.com/ablqk/tiny-go-projects/chapter-03/exercises/3_all_implementations/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo)

	// This produces nothing
	lgr.Debugf("Make the zero (%d) value useful.", 0)
}
