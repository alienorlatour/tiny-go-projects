package money

import (
	"reflect"
	"testing"
)

func TestApplyChangeRate(t *testing.T) {
	tt := map[string]struct {
		in             Amount
		rate           ExchangeRate
		targetCurrency Currency
		expected       Amount
	}{
		"Amount(1.52) * rate(1)": {
			in: Amount{
				quantity: Quantity{
					cents: 152,
					exp:   2,
				},
				currency: Currency{code: "TST", precision: 2},
			},
			rate:           1,
			targetCurrency: Currency{code: "TRG", precision: 4},
			expected: Amount{
				quantity: Quantity{
					cents: 15200,
					exp:   4,
				},
				currency: Currency{code: "TRG", precision: 4},
			},
		},
		"Amount(2.50) * rate(4)": {
			in: Amount{
				quantity: Quantity{
					cents: 250,
					exp:   2,
				}},
			rate:           4,
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Quantity{
					cents: 1000,
					exp:   2,
				},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		"Amount(4) * rate(2.5)": {
			in: Amount{
				quantity: Quantity{
					cents: 4,
					exp:   0,
				},
			},
			rate:           2.5,
			targetCurrency: Currency{code: "TRG", precision: 0},
			expected: Amount{
				quantity: Quantity{
					cents: 10,
					exp:   0,
				},
				currency: Currency{code: "TRG", precision: 0},
			},
		},
		"Amount(3.14) * rate(2.52678)": {
			in: Amount{
				quantity: Quantity{
					cents: 314,
					exp:   2,
				}},
			rate:           2.52678,
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Quantity{
					cents: 793,
					exp:   2,
				},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		"Amount(1.1) * rate(10)": {
			in: Amount{
				quantity: Quantity{
					cents: 11,
					exp:   1,
				}},
			rate:           10,
			targetCurrency: Currency{code: "TRG", precision: 1},
			expected: Amount{
				quantity: Quantity{
					cents: 110,
					exp:   1,
				},
				currency: Currency{code: "TRG", precision: 1},
			},
		},
		"Amount(1_000_000_000.01) * rate(2)": {
			in: Amount{
				quantity: Quantity{
					cents: 1_000_000_001,
					exp:   2,
				}},
			rate:           2,
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Quantity{
					cents: 2_000_000_002,
					exp:   2,
				},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		"Amount(265_413.87) * rate(5.05935e-5)": {
			in: Amount{
				quantity: Quantity{
					cents: 265_413_87,
					exp:   2,
				}},
			rate:           5.05935e-5,
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Quantity{
					cents: 13_42,
					exp:   2,
				},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		"Amount(265_413) * rate(1)": {
			in: Amount{
				quantity: Quantity{
					cents: 265_413,
					exp:   0,
				}},
			rate:           1,
			targetCurrency: Currency{code: "TRG", precision: 3},
			expected: Amount{
				quantity: Quantity{
					cents: 265_413_000,
					exp:   3,
				},
				currency: Currency{code: "TRG", precision: 3},
			},
		},
		"Amount(2) * rate(1.337)": {
			in: Amount{
				quantity: Quantity{
					cents: 2,
					exp:   0,
				}},
			rate:           1.337,
			targetCurrency: Currency{code: "TRG", precision: 5},
			expected: Amount{
				quantity: Quantity{
					cents: 267400,
					exp:   5,
				},
				currency: Currency{code: "TRG", precision: 5},
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := applyChangeRate(tc.in, tc.targetCurrency, tc.rate)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
