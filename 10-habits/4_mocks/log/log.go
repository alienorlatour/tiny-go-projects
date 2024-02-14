package log

import (
	"io"
	"log"
)

var logger log.Logger

func Set(writer io.Writer) {
	l := log.New(writer, "habit tracker", log.Ldate+log.Ltime)
	logger = *l
}
