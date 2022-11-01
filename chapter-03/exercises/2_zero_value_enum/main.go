package main

import (
	"fmt"

	"github.com/ablqk/tiny-go-projects/chapter-03/exercises/2_zero_value_enum/pocketlog"
)

func main() {
	l := pocketlog.Logger{}
	fmt.Printf("Logger: %#v\n", l)
}
