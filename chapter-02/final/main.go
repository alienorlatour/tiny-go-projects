package main

import (
	"os"
	"time"

	"tiny-go-projects/chapter-02/final/logger"
)

var lgr *logger.Logger

func main() {
	lgr.Info("Hallo, Welt")
	lgr.Error("Hello %s", "Susan")
	lgr.Debug("Hello %s", "Paul")

	lgr.Info("Hallo, %d %v", 2022, time.Now())
}

func init() {
	lgr = logger.New(logger.LevelInfo, logger.WithOutput(os.Stdout))
}
