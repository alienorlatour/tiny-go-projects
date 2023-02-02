package money

import (
	"reflect"
	"testing"
)

// TODO FIX me check amount with currencies and not only the Number
func TestApplyChangeRate(t *testing.T) {
	tt := map[string]struct {
		in             Amount
		rate           ExchangeRate
		targetCurrency Currency
		expected       Number
	}{
		"Number(1.52) * rate(1)": {
			in: Amount{
				number: Number{
					integerPart: 1,
					decimalPart: 52,
					precision:   2,
				},
				currency: Currency{},
			},
			rate:           1,
			targetCurrency: Currency{precision: 4},
			expected: Number{
				integerPart: 1,
				decimalPart: 5200,
				precision:   4,
			},
		},
		"Number(2.50) * rate(4)": {
			in: Amount{
				number: Number{
					integerPart: 2,
					decimalPart: 50,
					precision:   2,
				}},
			rate:           4,
			targetCurrency: Currency{precision: 2},
			expected: Number{
				integerPart: 10,
				decimalPart: 0,
				precision:   2,
			},
		},
		"Number(4) * rate(2.5)": {
			in: Amount{
				number: Number{
					integerPart: 4,
					decimalPart: 0,
					precision:   0,
				}},
			rate:           2.5,
			targetCurrency: Currency{precision: 0},
			expected: Number{
				integerPart: 10,
				decimalPart: 0,
				precision:   0,
			},
		},
		"Number(3.14) * rate(2.52678)": {
			in: Amount{
				number: Number{
					integerPart: 3,
					decimalPart: 14,
					precision:   2,
				}},
			rate:           2.52678,
			targetCurrency: Currency{precision: 2},
			expected: Number{
				integerPart: 7,
				decimalPart: 93,
				precision:   2,
			},
		},
		"Number(1.1) * rate(10)": {
			in: Amount{
				number: Number{
					integerPart: 1,
					decimalPart: 1,
					precision:   1,
				}},
			rate:           10,
			targetCurrency: Currency{precision: 1},
			expected: Number{
				integerPart: 11,
				decimalPart: 0,
				precision:   1,
			},
		},
		"Number(1_000_000_000.01) * rate(2)": {
			in: Amount{
				number: Number{
					integerPart: 1_000_000_000,
					decimalPart: 1,
					precision:   2,
				}},
			rate:           2,
			targetCurrency: Currency{precision: 2},
			expected: Number{
				integerPart: 2_000_000_000,
				decimalPart: 2,
				precision:   2,
			},
		},
		"Number(265_413.87) * rate(5.05935e-5)": {
			in: Amount{
				number: Number{
					integerPart: 265_413,
					decimalPart: 87,
					precision:   2,
				}},
			rate:           5.05935e-5,
			targetCurrency: Currency{precision: 2},
			expected: Number{
				integerPart: 13,
				decimalPart: 43,
				precision:   2,
			},
		},
		"Number(265_413) * rate(1)": {
			in: Amount{
				number: Number{
					integerPart: 265_413,
					decimalPart: 0,
					precision:   0,
				}},
			rate:           1,
			targetCurrency: Currency{precision: 3},
			expected: Number{
				integerPart: 265413,
				decimalPart: 0,
				precision:   3,
			},
		},
		"Number(2) * rate(1.337)": {
			in: Amount{
				number: Number{
					integerPart: 2,
					decimalPart: 0,
					precision:   0,
				}},
			rate:           1.337,
			targetCurrency: Currency{precision: 5},
			expected: Number{
				integerPart: 2,
				decimalPart: 67400,
				precision:   5,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := applyChangeRate(tc.in, tc.rate, tc.targetCurrency)
			if !reflect.DeepEqual(got.number, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
