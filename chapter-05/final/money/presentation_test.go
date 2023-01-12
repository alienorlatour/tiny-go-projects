package money_test

import (
	"errors"
	"testing"

	"github.com/ablqk/tiny-go-projects/chapter-05/layered/money"
)

func TestParseCurrencies(t *testing.T) {
	tt := map[string]struct {
		from            string
		to              string
		targetPrecision int
		validate        func(t *testing.T, gotFrom, gotTo money.Currency, err error)
	}{
		"USD to EUR": {
			from:            "USD",
			to:              "EUR",
			targetPrecision: 2,
			validate: func(t *testing.T, gotFrom, gotTo money.Currency, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}
				if gotFrom != money.NewCurrency("USD", 2, 0) {
					t.Errorf("expected USD, got %v", gotFrom)
				}
				if gotTo != money.NewCurrency("EUR", 2, 0) {
					t.Errorf("expected EUR, got %v", gotFrom)
				}
			},
		},
		"Unknown currency": {
			from:            "EUR",
			to:              "SUR", // Soviet Union Rubles, long gone.
			targetPrecision: 2,
			validate: func(t *testing.T, gotFrom, gotTo money.Currency, err error) {
				if !errors.Is(err, money.ErrUnknownCurrency) {
					t.Errorf("expected error %s, got %v", money.ErrUnknownCurrency, err)
				}
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			from, to, err := money.ParseCurrencies(tc.from, tc.to)
			tc.validate(t, from, to, err)
		})
	}
}
