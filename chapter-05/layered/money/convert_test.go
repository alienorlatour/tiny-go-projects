package money_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/ablqk/tiny-go-projects/chapter-05/layered/money"
)

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount          string
		from            string
		to              string
		targetPrecision int
		rateRepo        stubRate
		validate        func(t *testing.T, got string, err error)
	}{
		"34.98 USD to EUR": {
			amount:          "34.98",
			from:            "USD",
			to:              "EUR",
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 1.2564},
			validate: func(t *testing.T, got string, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}
				if got != "43.95" {
					t.Errorf("expected 53.06, got %q", got)
				}
			},
		},
		"Input amount is too large": {
			amount:          "34345982398459834.98",
			from:            "EUR",
			to:              "KRW",
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 1.5},
			validate: func(t *testing.T, got string, err error) {
				if !errors.Is(err, money.ErrInputTooLarge) {
					t.Errorf("expected error %s, got %v", money.ErrInputTooLarge, err)
				}
			},
		},
		"Input amount is too small": {
			amount:          "0.001",
			from:            "EUR",
			to:              "KRW",
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 1.5},
			validate: func(t *testing.T, got string, err error) {
				if !errors.Is(err, money.ErrInputTooSmall) {
					t.Errorf("expected error %s, got %v", money.ErrInputTooSmall, err)
				}
			},
		},
		"Output amount is too large": {
			amount:          "12345678901.23",
			from:            "EUR",
			to:              "IDR",
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 16_468.30},
			validate: func(t *testing.T, got string, err error) {
				if !errors.Is(err, money.ErrOutputTooLarge) {
					t.Errorf("expected error %s, got %v", money.ErrOutputTooLarge, err)
				}
			},
		},
		"Output amount is too small": {
			amount:          "150",
			from:            "IDR",
			to:              "EUR",
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 0.000060722722},
			validate: func(t *testing.T, got string, err error) {
				if !errors.Is(err, money.ErrOutputTooSmall) {
					t.Errorf("expected error %s, got %v", money.ErrOutputTooSmall, err)
				}
			},
		},
		"Unknown currency": {
			amount:          "10",
			from:            "EUR",
			to:              "SUR", // Soviet Union Rubles, long gone.
			targetPrecision: 2,
			rateRepo:        stubRate{err: fmt.Errorf("unknown currency")},
			validate: func(t *testing.T, got string, err error) {
				if !errors.Is(err, money.ErrUnknownChangeRate) {
					t.Errorf("expected error %s, got %v", money.ErrUnknownChangeRate, err)
				}
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(context.Background(), tc.amount, tc.from, tc.to, tc.rateRepo)
			tc.validate(t, got, err)
		})
	}
}

// stubRate is a very simple stub for the rateRepository.
type stubRate struct {
	rate money.ChangeRate
	err  error
}

// ExchangeRate implements the interface rateRepository with the same signature but fields are unused for tests purposes.
func (m stubRate) ExchangeRate(ctx context.Context, source, target money.Currency) (money.ChangeRate, error) {
	return m.rate, m.err
}
