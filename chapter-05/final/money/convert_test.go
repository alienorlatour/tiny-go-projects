package money_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/ablqk/tiny-go-projects/chapter-05/final/money"
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
			amount:          mustParseAmount(t, "34.98", "USD"),
			to:              mustParseCurrency(t, "EUR"),
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 1.2564},
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}
				expected := mustParseAmount(t, "43.95", "EUR")
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("expected %q, got %q", expected, got)
				}
			},
		},
		"Input amount is too large": {
			amount:          mustParseAmount(t, "34345982398459834.98", "EUR"),
			to:              mustParseCurrency(t, "KRW"),
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 1.5},
			validate: func(t *testing.T, got money.Amount, err error) {
				if !errors.Is(err, money.ErrTooLarge) {
					t.Errorf("expected error %s, got %v", money.ErrTooLarge, err)
				}
			},
		},
		"Input amount is too small": {
			amount:          mustParseAmount(t, "0.001", "EUR"),
			to:              mustParseCurrency(t, "KRW"),
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 1.5},
			validate: func(t *testing.T, got money.Amount, err error) {
				if !errors.Is(err, money.ErrTooSmall) {
					t.Errorf("expected error %s, got %v", money.ErrTooSmall, err)
				}
			},
		},
		"Output amount is too large": {
			amount:          mustParseAmount(t, "12345678901.23", "EUR"),
			to:              mustParseCurrency(t, "IDR"),
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 16_468.30},
			validate: func(t *testing.T, got money.Amount, err error) {
				if !errors.Is(err, money.ErrTooLarge) {
					t.Errorf("expected error %s, got %v", money.ErrTooLarge, err)
				}
			},
		},
		"Output amount is too small": {
			amount:          mustParseAmount(t, "150", "IDR"),
			to:              mustParseCurrency(t, "EUR"),
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 0.000060722722},
			validate: func(t *testing.T, got money.Amount, err error) {
				if !errors.Is(err, money.ErrTooSmall) {
					t.Errorf("expected error %s, got %v", money.ErrTooSmall, err)
				}
			},
		},
		"Unknown currency": {
			amount:          mustParseAmount(t, "10", "EUR"),
			to:              mustParseCurrency(t, "SUR"), // Soviet Union Rubles, long gone.
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
			got, err := money.Convert(tc.amount, tc.to, tc.rateRepo)
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
func (m stubRate) FetchExchangeRate(_, _ money.Currency) (money.ExchangeRate, error) {
	return m.rate, m.err
}

func mustParseAmount(t *testing.T, value string, code string) money.Amount {
	n, err := money.ParseNumber(value)
	if err != nil {
		t.Fatalf("cannot parse value: %s", value)
	}

	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("cannot parse currency code: %s", code)
	}

	amount, err := money.NewAmount(n, currency)
	if err != nil {
		t.Fatalf("cannot create amount with value %s and currency code %s", amount, code)
	}

	return amount
}

func mustParseCurrency(t *testing.T, code string) money.Currency {
	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("cannot parse currency %s code", code)
	}

	return currency
}
