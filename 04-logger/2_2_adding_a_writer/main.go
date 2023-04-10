package main

import (
	"os"

	"learngo-pockets/logger/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, os.Stdout)

	// This produces nothing
	lgr.Debugf("Make the zero (%d) value useful.", 0)
}
