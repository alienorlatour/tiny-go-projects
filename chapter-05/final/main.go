package main

import (
	"flag"
	"fmt"
	"time"
)

// Usage: calconvert -from gregorian -to hijri -date “5 February 2022”

func main() {
	from := flag.String("from", "gregorian", "source calendar")
	to := flag.String("to", "gregorian", "target calendar")
	date := flag.String("date", time.Now().Format("2006-01-02"), "date to convert")

	flag.Parse()

	d := calendars.New(from).Parse(date)
	fmt.Println(calendars.New(to).Format(d))
}
