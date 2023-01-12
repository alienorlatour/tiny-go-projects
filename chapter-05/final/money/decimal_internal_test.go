package money

import (
	"errors"
	"testing"
)

func TestParseNumber(t *testing.T) {
	tt := map[string]struct {
		amount   string
		expected Number
		err      error
	}{
		"2 decimal digits": {
			amount:   "1.52",
			expected: Number{integerPart: 152, precision: 2},
			err:      nil,
		},
		"no decimal digits": {
			amount:   "1",
			expected: Number{integerPart: 1, precision: 0},
			err:      nil,
		},
		"suffix 0 as decimal digits": {
			amount:   "1.50",
			expected: Number{integerPart: 150, precision: 2},
			err:      nil,
		},
		"prefix 0 as decimal digits": {
			amount:   "1.02",
			expected: Number{integerPart: 102, precision: 2},
			err:      nil,
		},
		"with apostrophes for readability": {
			amount: "12'152.03",
			// expected: Number{integerPart: 121523, precision: 2}, // for future implementations
			err: errInvalidInteger,
		},
		"with underscores for readability": {
			amount: "12_152.03",
			// expected: Number{integerPart: 1215203, precision: 2}, // for future implementations
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
		n        Number
		expected string
	}{
		"15.2": {
			n: Number{
				integerPart: 152,
				precision:   1,
			},
			expected: "15.2",
		},
		"15.02": {
			n: Number{
				integerPart: 1502,
				precision:   2,
			},
			expected: "15.02",
		},
		"15.0200": {
			n: Number{
				integerPart: 150200,
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
		in              Number
		rate            ChangeRate
		targetPrecision int
		expected        Number
	}{
		"Number(1.52) * rate(1)": {
			in: Number{
				integerPart: 152,
				precision:   2,
			},
			rate:            1,
			targetPrecision: 4,
			expected: Number{
				integerPart: 15200,
				precision:   4,
			},
		},
		"Number(2.50) * rate(4)": {
			in: Number{
				integerPart: 250,
				precision:   2,
			},
			rate:            4,
			targetPrecision: 2,
			expected: Number{
				integerPart: 1000,
				precision:   2,
			},
		},
		"Number(4) * rate(2.5)": {
			in: Number{
				integerPart: 4,
				precision:   0,
			},
			rate:            2.5,
			targetPrecision: 0,
			expected: Number{
				integerPart: 10,
				precision:   0,
			},
		},
		"Number(3.14) * rate(2.52678)": {
			in: Number{
				integerPart: 314,
				precision:   2,
			},
			rate:            2.52678,
			targetPrecision: 2,
			expected: Number{
				integerPart: 793,
				precision:   2,
			},
		},
		"Number(1.1) * rate(10)": {
			in: Number{
				integerPart: 11,
				precision:   1,
			},
			rate:            10,
			targetPrecision: 1,
			expected: Number{
				integerPart: 110,
				precision:   1,
			},
		},
		"Number(1_000_000_000.01) * rate(2)": {
			in: Number{
				integerPart: 100_000_000_001,
				precision:   2,
			},
			rate:            2,
			targetPrecision: 2,
			expected: Number{
				integerPart: 200_000_000_002,
				precision:   2,
			},
		},
		"Number(265_413.87) * rate(5.05935e-5)": {
			in: Number{
				integerPart: 26_541_387,
				precision:   2,
			},
			rate:            5.05935e-5,
			targetPrecision: 2,
			expected: Number{
				integerPart: 1343,
				precision:   2,
			},
		},
		"Number(265_413) * rate(1)": {
			in: Number{
				integerPart: 265_413,
				precision:   0,
			},
			rate:            1,
			targetPrecision: 3,
			expected: Number{
				integerPart: 265413000,
				precision:   3,
			},
		},
		"Number(2) * rate(1.337)": {
			in: Number{
				integerPart: 2,
				precision:   0,
			},
			rate:            1.337,
			targetPrecision: 5,
			expected: Number{
				integerPart: 267400,
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
