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
			expected: number{integerPart: 1, decimalPart: 52, precision: 2},
			err:      nil,
		},
		"no decimal digits": {
			amount:   "1",
			expected: number{integerPart: 1, decimalPart: 0, precision: 0},
			err:      nil,
		},
		"suffix 0 as decimal digits": {
			amount:   "1.50",
			expected: number{integerPart: 1, decimalPart: 50, precision: 2},
			err:      nil,
		},
		"prefix 0 as decimal digits": {
			amount:   "1.02",
			expected: number{integerPart: 1, decimalPart: 2, precision: 2},
			err:      nil,
		},
		"invalid decimal part": {
			amount: "65.pocket",
			err:    errInvalidDecimal,
		},
		"with apostrophes for readability": {
			amount: "12'152.03",
			// expected: number{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
			err: errInvalidInteger,
		},
		"with underscores for readability": {
			amount: "12_152.03",
			// expected: number{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
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
				precision:   1,
			},
			expected: "15.2",
		},
		"15.02": {
			n: number{
				integerPart: 15,
				decimalPart: 2,
				precision:   2,
			},
			expected: "15.02",
		},
		"15.0200": {
			n: number{
				integerPart: 15,
				decimalPart: 200,
				precision:   4,
			},
			expected: "15.0200",
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
		in              number
		rate            ChangeRate
		targetPrecision int
		expected        number
	}{
		"number(1.52) * rate(1)": {
			in: number{
				integerPart: 1,
				decimalPart: 52,
				precision:   2,
			},
			rate:            1,
			targetPrecision: 4,
			expected: number{
				integerPart: 1,
				decimalPart: 5200,
				precision:   4,
			},
		},
		"number(2.50) * rate(4)": {
			in: number{
				integerPart: 2,
				decimalPart: 50,
				precision:   2,
			},
			rate:            4,
			targetPrecision: 2,
			expected: number{
				integerPart: 10,
				decimalPart: 0,
				precision:   2,
			},
		},
		"number(4) * rate(2.5)": {
			in: number{
				integerPart: 4,
				decimalPart: 0,
				precision:   0,
			},
			rate:            2.5,
			targetPrecision: 0,
			expected: number{
				integerPart: 10,
				decimalPart: 0,
				precision:   0,
			},
		},
		"number(3.14) * rate(2.52678)": {
			in: number{
				integerPart: 3,
				decimalPart: 14,
				precision:   2,
			},
			rate:            2.52678,
			targetPrecision: 2,
			expected: number{
				integerPart: 7,
				decimalPart: 93,
				precision:   2,
			},
		},
		"number(1.1) * rate(10)": {
			in: number{
				integerPart: 1,
				decimalPart: 1,
				precision:   1,
			},
			rate:            10,
			targetPrecision: 1,
			expected: number{
				integerPart: 11,
				decimalPart: 0,
				precision:   1,
			},
		},
		"number(1_000_000_000.01) * rate(2)": {
			in: number{
				integerPart: 1_000_000_000,
				decimalPart: 1,
				precision:   2,
			},
			rate:            2,
			targetPrecision: 2,
			expected: number{
				integerPart: 2_000_000_000,
				decimalPart: 2,
				precision:   2,
			},
		},
		"number(265_413.87) * rate(5.05935e-5)": {
			in: number{
				integerPart: 265_413,
				decimalPart: 87,
				precision:   2,
			},
			rate:            5.05935e-5,
			targetPrecision: 2,
			expected: number{
				integerPart: 13,
				decimalPart: 43,
				precision:   2,
			},
		},
		"number(265_413) * rate(1)": {
			in: number{
				integerPart: 265_413,
				decimalPart: 0,
				precision:   0,
			},
			rate:            1,
			targetPrecision: 3,
			expected: number{
				integerPart: 265413,
				decimalPart: 0,
				precision:   3,
			},
		},
		"number(2) * rate(1.337)": {
			in: number{
				integerPart: 2,
				decimalPart: 0,
				precision:   0,
			},
			rate:            1.337,
			targetPrecision: 5,
			expected: number{
				integerPart: 2,
				decimalPart: 67400,
				precision:   5,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.in.applyChangeRate(tc.rate, tc.targetPrecision)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
