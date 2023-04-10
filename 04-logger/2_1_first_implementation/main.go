package main

import (
	"learngo-pockets/logger/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo)

	// This produces nothing
	lgr.Debugf("Make the zero (%d) value useful.", 0)
}
