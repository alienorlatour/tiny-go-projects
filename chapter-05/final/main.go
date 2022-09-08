package main

import (
	"flag"
	"fmt"
)

// Usage: change -from USD -to EUR 34

func main() {
	from := flag.String("from", "", "source currency")
	to := flag.String("to", "EUR", "target currency, default to euros")

	flag.Parse()

	amount := flag.Arg(0)

	fmt.Printf("convert %s %s to %s\n", amount, *from, *to)
}
