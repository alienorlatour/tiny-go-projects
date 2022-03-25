package main

type debugger interface {
	Debug(format string, args ...any)
}

type infoer interface {
	Info(format string, args ...any)
}

type errorer interface {
	Error(format string, args ...any)
}
