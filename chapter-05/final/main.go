package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-05/final/money"
)

// Usage: change -from USD -to EUR -precision 2 34.98

func main() {
	// read currencies from the input
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")
	targetPrecision := flag.Int("precision", 0, "target precision")

	flag.Parse()

	if *from == "" {
		flag.Usage()
		os.Exit(1)
	}

	// read the amount to convert from the command
	amount := flag.Arg(0)
	convertedAmount, err := money.Convert(amount, *from, *to, *targetPrecision)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to convert %q %q to %q: %s.", amount, *from, *to, err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s %s = %s %s\n", amount, *from, convertedAmount, *to)
}
