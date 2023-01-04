package money_test

import (
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
		"34.98 EUR to KRW": {
			amount:          "34345982398459834.98",
			from:            "EUR",
			to:              "KRW",
			targetPrecision: 2,
			rateRepo:        stubRate{rate: 1.5},
			validate: func(t *testing.T, got string, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}
				if got != "43.95" {
					t.Errorf("expected 53.06, got %q", got)
				}
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.from, tc.to, tc.rateRepo)
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
func (m stubRate) ExchangeRate(source, target money.Currency) (money.ChangeRate, error) {
	return m.rate, m.err
}
