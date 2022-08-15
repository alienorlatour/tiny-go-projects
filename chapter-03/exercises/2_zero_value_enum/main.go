package main

import (
	"fmt"

	"tiny-go-projects/chapter03/exercises/2_zero_value_enum/pocketlog"
)

func main() {
	l := pocketlog.Logger{}
	fmt.Printf("Logger: %#v\n", l)
}
