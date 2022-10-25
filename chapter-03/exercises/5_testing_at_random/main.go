package main

import (
	"os"
	"time"

	"tiny-go-projects/chapter03/exercises/5_testing_at_random/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(os.Stdout))

	lgr.Infof("Hallo, Welt")
	lgr.Errorf("Hello %s", "Susan")
	lgr.Debugf("Hello %s", "Paul")

	lgr.Infof("Hallo, %d %v", 2022, time.Now())
}
