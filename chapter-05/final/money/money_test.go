package money_test

import (
	"github.com/ablqk/tiny-go-projects/chapter-05/final/money"
	"testing"
)

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount          string
		from            string
		to              string
		targetPrecision int
		validate        func(t *testing.T, got string, err error)
	}{
		"34.98 USD to EUR": {
			amount:          "34.98",
			from:            "USD",
			to:              "EUR",
			targetPrecision: 2,
			validate: func(t *testing.T, got string, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}
				if got != "53.00" {
					t.Errorf("expected 53.06, got %q", got)
				}
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.from, tc.to, tc.targetPrecision)
			tc.validate(t, got, err)
		})
	}
}
