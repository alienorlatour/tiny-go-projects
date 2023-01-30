package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-05/layered/ecbank"
	"github.com/ablqk/tiny-go-projects/chapter-05/layered/money"
)

// Usage: change -from USD -to EUR 34.98

func main() {
	// read currencies from the input
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency")

	// parse flags
	flag.Parse()

	// validate inputs
	if err := validateInputs(*from, *to, flag.NArg()); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		flag.Usage()
		os.Exit(1)
	}

	// create the exchange rate we want to use
	exchangeRate := ecbank.New(ecbank.Host)

	// read the number to convert from the command
	value := flag.Arg(0)
	if value == "" {
		_, _ = fmt.Fprintln(os.Stderr, "missing the amount to convert")
		os.Exit(1)
	}

	// parse the source currency
	fromCurrency, err := money.ParseCurrency(*from)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse currency %q to %q: %s.\n", *from, *to, err.Error())
		os.Exit(1)
	}

	// parse the target currency
	toCurrency, err := money.ParseCurrency(*to)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse currency %q to %q: %s.\n", *from, *to, err.Error())
		os.Exit(1)
	}

	// parse into a number
	number, err := money.ParseNumber(value)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse value %q: %s.\n", value, err.Error())
		os.Exit(1)
	}

	// transform value into an amount with its currency
	amount := money.NewAmount(number, fromCurrency)

	ctx := context.Background()

	// convert the amount from the source currency to the target with the current exchange rate
	convertedAmount, err := money.Convert(ctx, amount, toCurrency, exchangeRate)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to convert %q %q to %q: %s.\n", value, *from, *to, err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s %s = %s\n", value, *from, convertedAmount.String())
}

// validateInputs returns an error if one of the input is missing.
func validateInputs(from, to string, argc int) error {
	if from == "" {
		return errors.New("missing input currency\n")
	}

	if to == "" {
		// will never happen, flag has a default value
		return errors.New("missing output currency\n")
	}

	if argc != 1 {
		return errors.New("invalid number of arguments, expecting only the amount to convert\n")
	}

	return nil
}
