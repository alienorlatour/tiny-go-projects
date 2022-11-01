package main

import (
	"os"
	"time"

	"github.com/ablqk/tiny-go-projects/chapter-03/exercises/8_avoiding_long_messages/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(os.Stdout))

	lgr.Infof("Hallo, Welt")
	lgr.Errorf("Hello %s", "Susan")
	lgr.Debugf("Hello %s", "Paul")

	lgr.Infof("Hallo, %d %v", 2022, time.Now())
}
