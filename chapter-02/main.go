package main

import "tiny-go-projects/chapter-02/logger"

var lgr *logger.Logger

func main() {
	lgr.Info("Hallo, Welt")
}

func init() {
	lgr = logger.New(logger.LevelInfo)
}
