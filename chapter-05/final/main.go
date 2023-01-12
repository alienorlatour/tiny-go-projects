package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-05/layered/money"
	"github.com/ablqk/tiny-go-projects/chapter-05/layered/repository"
)

const ecbRepoURL = "https://www.ecb.europa.eu/"

// Usage: change -from USD -to EUR 34.98

func main() {
	// read currencies from the input
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")

	flag.Parse()

	if *from == "" {
		flag.Usage()
		os.Exit(1)
	}

	// create the repository we want to use
	changeRepo := repository.New(ecbRepoURL)

	ctx := context.Background()

	// read the amount to convert from the command
	amount := flag.Arg(0)

	number, err := money.ParseAmount(amount)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse amount %q: %s.\n", amount, err.Error())
		os.Exit(1)
	}

	fromCurrency, toCurrency, err := money.ParseCurrencies(*from, *to)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse currency %q to %q: %s.\n", *from, *to, err.Error())
		os.Exit(1)
	}

	convertedAmount, err := money.Convert(ctx, number, fromCurrency, toCurrency, changeRepo)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to convert %q %q to %q: %s.\n", amount, *from, *to, err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s %s = %s %s\n", amount, *from, convertedAmount, *to)
}
