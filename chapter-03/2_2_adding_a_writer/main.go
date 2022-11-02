package main

import (
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-03/2_2_adding_a_writer/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, os.Stdout)

	// This produces nothing
	lgr.Debugf("Make the zero (%d) value useful.", 0)
}
