package main

import (
	"learngo-pockets/logger/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelDebug)

	// This produces nothing
	lgr.Debugf("Make the zero (%d) value useful.", 0)
}
