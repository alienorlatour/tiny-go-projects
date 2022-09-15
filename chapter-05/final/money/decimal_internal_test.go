package money

import (
	"errors"
	"testing"
)

func TestParseNumber(t *testing.T) {
	tt := map[string]struct {
		amount   string
		expected number
		err      error
	}{
		"2 decimal digits": {
			amount:   "1.52",
			expected: number{integerPart: 1, decimalPart: 52, toUnit: 2},
			err:      nil,
		},
		"no decimal digits": {
			amount:   "1",
			expected: number{integerPart: 1, decimalPart: 0, toUnit: 0},
			err:      nil,
		},
		"suffix 0 as decimal digits": {
			amount:   "1.50",
			expected: number{integerPart: 1, decimalPart: 50, toUnit: 2},
			err:      nil,
		},
		"prefix 0 as decimal digits": {
			amount:   "1.02",
			expected: number{integerPart: 1, decimalPart: 2, toUnit: 2},
			err:      nil,
		},
		"invalid decimal part": {
			amount: "65.pocket",
			err:    errInvalidDecimal,
		},
		"with apostrophes for readability": {
			amount: "12'152.03",
			//expected: number{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
			err: errInvalidInteger,
		},
		"with underscores for readability": {
			amount: "12_152.03",
			//expected: number{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
			err: errInvalidInteger,
		},
		"NaN": {
			amount: "ten",
			err:    errInvalidInteger,
		},
		"empty string": {
			amount: "",
			err:    errInvalidInteger,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := parseNumber(tc.amount)
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %q, got %q", tc.err.Error(), err.Error())
			}
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestNumberString(t *testing.T) {
	tt := map[string]struct {
		n        number
		expected string
	}{
		"15.2": {
			n: number{
				integerPart: 15,
				decimalPart: 2,
				toUnit:      1,
			},
			expected: "15.2",
		},
		"15.02": {
			n: number{
				integerPart: 15,
				decimalPart: 2,
				toUnit:      2,
			},
			expected: "15.02",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.n.String()
			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}

func TestNumberApplyChangeRate(t *testing.T) {
	tt := map[string]struct {
		in       number
		rate     changeRate
		expected number
	}{
		"rate is 1": {
			in: number{
				integerPart: 1,
				decimalPart: 52,
				toUnit:      2,
			},
			rate: 1,
			expected: number{
				integerPart: 1,
				decimalPart: 52,
				toUnit:      2,
			},
		},
		"2.50 * 4": {
			in: number{
				integerPart: 2,
				decimalPart: 50,
				toUnit:      2,
			},
			rate: 4,
			expected: number{
				integerPart: 10,
				decimalPart: 0,
				toUnit:      2,
			},
		},
		"4 * 2.5": {
			in: number{
				integerPart: 4,
				decimalPart: 0,
				toUnit:      0,
			},
			rate: 2.5,
			expected: number{
				integerPart: 10,
				decimalPart: 0,
				toUnit:      0,
			},
		},
		"3.14 * 2.52678": {
			in: number{
				integerPart: 3,
				decimalPart: 14,
				toUnit:      2,
			},
			rate: 2.52678,
			expected: number{
				integerPart: 7,
				decimalPart: 93,
				toUnit:      2,
			},
		},
		"1.1 * 10": {
			in: number{
				integerPart: 1,
				decimalPart: 1,
				toUnit:      1,
			},
			rate: 10,
			expected: number{
				integerPart: 11,
				decimalPart: 0,
				toUnit:      1,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.in.applyChangeRate(tc.rate)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
