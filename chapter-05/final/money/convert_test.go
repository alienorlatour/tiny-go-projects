package money_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/ablqk/tiny-go-projects/chapter-05/layered/money"
)

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount          money.Amount
		to              money.Currency
		targetPrecision int
		rateRepo        stubRate
		validate        func(t *testing.T, got money.Amount, err error)
	}{
		"34.98 USD to EUR": {
			amount:          money.NewAmountHelper("34.98", "USD"),
			to:              money.NewCurrency("EUR", 2, 0),
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 1.2564},
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}
				expected := money.NewAmountHelper("43.95", "EUR")
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("expected %q, got %q", expected, got)
				}
			},
		},
		"Input amount is too large": {
			amount:          money.NewAmountHelper("34345982398459834.98", "EUR"),
			to:              money.NewCurrency("KRW", 2, 0),
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 1.5},
			validate: func(t *testing.T, got money.Amount, err error) {
				if !errors.Is(err, money.ErrTooLarge) {
					t.Errorf("expected error %s, got %v", money.ErrTooLarge, err)
				}
			},
		},
		"Input amount is too small": {
			amount:          money.NewAmountHelper("0.001", "EUR"),
			to:              money.NewCurrency("KRW", 2, 0),
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 1.5},
			validate: func(t *testing.T, got money.Amount, err error) {
				if !errors.Is(err, money.ErrTooSmall) {
					t.Errorf("expected error %s, got %v", money.ErrTooSmall, err)
				}
			},
		},
		"Output amount is too large": {
			amount:          money.NewAmountHelper("12345678901.23", "EUR"),
			to:              money.NewCurrency("IDR", 2, 0),
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 16_468.30},
			validate: func(t *testing.T, got money.Amount, err error) {
				if !errors.Is(err, money.ErrTooLarge) {
					t.Errorf("expected error %s, got %v", money.ErrTooLarge, err)
				}
			},
		},
		//// TODO FIX me
		//"Output amount is too small": {
		//	amount:          money.NewAmountHelper("150", "IDR"),
		//	to:              money.NewCurrency("EUR", 2, 0),
		//	targetPrecision: 2,
		//	rateRepo:        stubRate{rate: 0.000060722722},
		//	validate: func(t *testing.T, got money.Amount, err error) {
		//		if !errors.Is(err, money.ErrOutputTooSmall) {
		//			t.Errorf("expected error %s, got %v", money.ErrOutputTooSmall, err)
		//		}
		//	},
		//},
		"Unknown currency": {
			amount:          money.NewAmountHelper("10", "EUR"),
			to:              money.NewCurrency("SUR", 2, 0), // Soviet Union Rubles, long gone.
			targetPrecision: 2,
			rateRepo:        stubRate{err: fmt.Errorf("unknown currency")},
			validate: func(t *testing.T, got money.Amount, err error) {
				if !errors.Is(err, money.ErrGettingChangeRate) {
					t.Errorf("expected error %s, got %v", money.ErrGettingChangeRate, err)
				}
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(context.Background(), tc.amount, tc.to, tc.rateRepo)
			tc.validate(t, got, err)
		})
	}
}

// stubRate is a very simple stub for the exchangeRates.
type stubRate struct {
	rate money.ExchangeRate
	err  error
}

// ExchangeRate implements the interface exchangeRates with the same signature but fields are unused for tests purposes.
func (m stubRate) FetchExchangeRate(ctx context.Context, source, target money.Currency) (money.ExchangeRate, error) {
	return m.rate, m.err
}
