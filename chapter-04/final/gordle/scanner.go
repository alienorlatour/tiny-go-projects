package gordle

// scanner is a local interface that is implemented by bufio.Scanner. And also by testScanner. How fortunate.
type scanner interface {
	// Scan returns whether there is still data to read.
	Scan() bool
	// Text returns the data that has been read since the last call to Scan.
	Text() string
	// Err returns the error that was received since the last call to Scan.
	Err() error
}
