package money_test

import (
	"reflect"
	"testing"

	"github.com/ablqk/tiny-go-projects/chapter-05/2_represent/money"
)

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount   money.Amount
		to       money.Currency
		validate func(t *testing.T, got money.Amount, err error)
	}{
		"34.98 USD to EUR": {
			amount: mustParseAmount(t, "34.98", "USD"),
			to:     mustParseCurrency(t, "EUR"),
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
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to, nil)
			tc.validate(t, got, err)
		})
	}
}
