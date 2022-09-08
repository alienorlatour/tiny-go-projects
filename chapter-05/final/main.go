package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-05/final/money"
)

// Usage: change -from USD -to EUR 34.98

func main() {
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")

	flag.Parse()

	amount := flag.Arg(0)

	if *from == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("%s %s = %s %s\n", amount, *from, money.Convert(amount, *from, *to), *to)
}
