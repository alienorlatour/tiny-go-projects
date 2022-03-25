package main

import "github.com/ablqk/tiny-go-projects/chapter-02/logger"

type debugger interface {
	Debug(format string, args ...any)
}

type infoer interface {
	Info(format string, args ...any)
}

type errorer interface {
	Error(format string, args ...any)
}

func main() {
	l := logger.New(logger.LevelInfo)
	l.Info("Hello, world")
}
