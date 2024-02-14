package hlog

import (
	"io"
	"log"
)

// logger is the common instance.
var logger log.Logger

func Set(writer io.Writer) {
	l := log.New(writer, "habit tracker", log.Ldate+log.Ltime)
	logger = *l // FIXME Assignment copies a lock value to 'logger': type 'log.Logger' contains 'sync.Mutex' which is 'sync.Locker'
}
