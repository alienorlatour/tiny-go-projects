package money_test

import (
	"testing"

	"github.com/ablqk/tiny-go-projects/chapter-05/final/money"
)

func TestAmountString(t *testing.T) {
	tt := map[string]struct {
		amount   money.Amount
		expected string
	}{
		"15.2 EUR": {
			amount:   mustParseAmount(t, "15.2", "EUR"),
			expected: "15.2 EUR",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.amount.String()
			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}
