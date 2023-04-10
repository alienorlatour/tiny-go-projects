package money

import (
	"errors"
	"reflect"
	"testing"
)

func TestApplyExchangeRate(t *testing.T) {
	tt := map[string]struct {
		in             Amount
		rate           ExchangeRate
		targetCurrency Currency
		expected       Amount
		err            error
	}{
		"Amount(1.52) * rate(1)": {
			in: Amount{
				quantity: Decimal{
					subunits:  152,
					precision: 2,
				},
				currency: Currency{code: "TST", precision: 2},
			},
			rate:           ExchangeRate{subunits: 1, precision: 0},
			targetCurrency: Currency{code: "TRG", precision: 4},
			expected: Amount{
				quantity: Decimal{
					subunits:  15200,
					precision: 4,
				},
				currency: Currency{code: "TRG", precision: 4},
			},
		},
		"Amount(2.50) * rate(4)": {
			in: Amount{
				quantity: Decimal{
					subunits:  250,
					precision: 2,
				}},
			rate:           ExchangeRate{subunits: 4, precision: 0},
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Decimal{
					subunits:  1000,
					precision: 2,
				},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		"Amount(4) * rate(2.5)": {
			in: Amount{
				quantity: Decimal{
					subunits:  4,
					precision: 0,
				},
			},
			rate:           ExchangeRate{subunits: 25, precision: 1},
			targetCurrency: Currency{code: "TRG", precision: 0},
			expected: Amount{
				quantity: Decimal{
					subunits:  10,
					precision: 0,
				},
				currency: Currency{code: "TRG", precision: 0},
			},
		},
		"Amount(3.14) * rate(2.52678)": {
			in: Amount{
				quantity: Decimal{
					subunits:  314,
					precision: 2,
				}},
			rate:           ExchangeRate{subunits: 252678, precision: 5},
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Decimal{
					subunits:  793,
					precision: 2,
				},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		"Amount(1.1) * rate(10)": {
			in: Amount{
				quantity: Decimal{
					subunits:  11,
					precision: 1,
				}},
			rate:           ExchangeRate{subunits: 10, precision: 0},
			targetCurrency: Currency{code: "TRG", precision: 1},
			expected: Amount{
				quantity: Decimal{
					subunits:  110,
					precision: 1,
				},
				currency: Currency{code: "TRG", precision: 1},
			},
		},
		"Amount(1_000_000_000.01) * rate(2)": {
			in: Amount{
				quantity: Decimal{
					subunits:  1_000_000_001,
					precision: 2,
				}},
			rate:           ExchangeRate{subunits: 2, precision: 0},
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Decimal{
					subunits:  2_000_000_002,
					precision: 2,
				},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		"Amount(265_413.87) * rate(5.05935e-5)": {
			in: Amount{
				quantity: Decimal{
					subunits:  265_413_87,
					precision: 2,
				}},
			rate:           ExchangeRate{subunits: 505935, precision: 10},
			targetCurrency: Currency{code: "TRG", precision: 2},
			expected: Amount{
				quantity: Decimal{
					subunits:  13_42,
					precision: 2,
				},
				currency: Currency{code: "TRG", precision: 2},
			},
		},
		"Amount(265_413) * rate(1)": {
			in: Amount{
				quantity: Decimal{
					subunits:  265_413,
					precision: 0,
				}},
			rate:           ExchangeRate{subunits: 1, precision: 0},
			targetCurrency: Currency{code: "TRG", precision: 3},
			expected: Amount{
				quantity: Decimal{
					subunits:  265_413_000,
					precision: 3,
				},
				currency: Currency{code: "TRG", precision: 3},
			},
		},
		"Amount(2) * rate(1.337)": {
			in: Amount{
				quantity: Decimal{
					subunits:  2,
					precision: 0,
				}},
			rate:           ExchangeRate{subunits: 1337, precision: 3},
			targetCurrency: Currency{code: "TRG", precision: 5},
			expected: Amount{
				quantity: Decimal{
					subunits:  267400,
					precision: 5,
				},
				currency: Currency{code: "TRG", precision: 5},
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := applyExchangeRate(tc.in, tc.targetCurrency, tc.rate)
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tt := map[string]struct {
		decimal Decimal
		rate    ExchangeRate
		want    Decimal
	}{
		"10.0 * 1.2 = 12": {
			decimal: Decimal{100, 1},
			rate:    ExchangeRate{subunits: 12, precision: 1},
			want:    Decimal{12, 0},
		},
		"10 * 2 = 20": {
			decimal: Decimal{10, 0},
			rate:    ExchangeRate{subunits: 2, precision: 0},
			want:    Decimal{20, 0},
		},
		"10 * 1e15": {
			decimal: Decimal{10, 0},
			rate:    ExchangeRate{subunits: 1_000_000_000_000_000, precision: 0},
			want:    Decimal{subunits: 10_000_000_000_000_000, precision: 0},
		},
		"1698.188 * 2.198677818 = 3733.768286393784": {
			decimal: Decimal{1698188, 3},
			rate:    ExchangeRate{subunits: 2198677818, precision: 9},
			want:    Decimal{3733768286393784, 12},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := multiply(tc.decimal, tc.rate)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}
