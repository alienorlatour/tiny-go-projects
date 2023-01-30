package money

import "testing"

func TestAmountString(t *testing.T) {
	tt := map[string]struct {
		amount   Amount
		expected string
	}{
		"15.2 EUR": {
			amount: Amount{
				number: Number{
					integerPart: 15,
					decimalPart: 2,
					precision:   1,
				},
				currency: NewCurrency("EUR", 2, 1),
			},

			expected: "15.2 EUR",
		},
		"missing Currency": {
			amount: Amount{
				number: Number{
					integerPart: 15,
					decimalPart: 2,
					precision:   1,
				},
			},

			expected: "15.2 ",
		},
		"missing Number": {
			amount:   Amount{currency: NewCurrency("EUR", 2, 1)},
			expected: "0.0 EUR",
		},
		"missing Number and Currency": {
			amount:   Amount{},
			expected: "0.0 ",
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
