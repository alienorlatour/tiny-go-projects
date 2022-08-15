package main

import (
	"os"
	"time"

	"tiny-go-projects/chapter03/exercises/6_logging_level/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(os.Stdout))

	lgr.Info("Hallo, Welt")
	lgr.Error("Hello %s", "Susan")
	lgr.Debug("Hello %s", "Paul")

	lgr.Info("Hallo, %d %v", 2022, time.Now())
}
