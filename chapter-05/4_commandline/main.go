package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-05/4_commandline/money"
)

func main() {
	// read currencies from the input
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")

	// parse flags
	flag.Parse()

	// parse the source currency
	fromCurrency, err := money.ParseCurrency(*from)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse currency %q: %s.\n", *from, err.Error())
		os.Exit(1)
	}

	// parse the target currency
	toCurrency, err := money.ParseCurrency(*to)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse currency %q: %s.\n", *to, err.Error())
		os.Exit(1)
	}

	// read the argument
	value := flag.Arg(0)
	if value == "" {
		_, _ = fmt.Fprintln(os.Stderr, "missing amount to convert")
		flag.Usage()
		os.Exit(1)
	}

	// parse into a number
	number, err := money.ParseNumber(value)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse value %q: %s.\n", value, err.Error())
		os.Exit(1)
	}

	// transform value into an amount with its currency
	amount, err := money.NewAmount(number, fromCurrency)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(amount, toCurrency)
}
