package main

import (
	"github.com/ablqk/tiny-go-projects/chapter-03/1_3_new/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// This produces nothing
	l.Debugf("hello")
}
