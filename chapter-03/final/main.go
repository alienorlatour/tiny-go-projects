package main

import (
	"os"
	"time"

	"tiny-go-projects/chapter03/final/pocketlog"
)

var lgr *pocketlog.Logger

func main() {
	lgr.Info("Hallo, Welt")
	lgr.Error("Hello %s", "Susan")
	lgr.Debug("Hello %s", "Paul")

	lgr.Info("Hallo, %d %v", 2022, time.Now())
}

func init() {
	lgr = pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(os.Stdout))
}
