package main

import "github.com/ablqk/tiny-go-projects/chapter-02/logger"

var lgr *logger.Logger

func init() {
	lgr = logger.New(logger.LevelInfo)
}

func main() {
	lgr.Info("Hallo, Welt")
}
