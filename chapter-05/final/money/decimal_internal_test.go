package money

import (
	"errors"
	"testing"
)

func TestParseAmount(t *testing.T) {
	tt := map[string]struct {
		amount   string
		expected Number
		err      error
	}{
		"2 decimal digits": {
			amount:   "1.52",
			expected: Number{integerPart: 1, decimalPart: 52, precision: 2},
			err:      nil,
		},
		"no decimal digits": {
			amount:   "1",
			expected: Number{integerPart: 1, decimalPart: 0, precision: 0},
			err:      nil,
		},
		"suffix 0 as decimal digits": {
			amount:   "1.50",
			expected: Number{integerPart: 1, decimalPart: 50, precision: 2},
			err:      nil,
		},
		"prefix 0 as decimal digits": {
			amount:   "1.02",
			expected: Number{integerPart: 1, decimalPart: 2, precision: 2},
			err:      nil,
		},
		"invalid decimal part": {
			amount: "65.pocket",
			err:    ErrInvalidDecimal,
		},
		"with apostrophes for readability": {
			amount: "12'152.03",
			// expected: Number{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
			err: ErrInvalidInteger,
		},
		"with underscores for readability": {
			amount: "12_152.03",
			// expected: Number{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
			err: ErrInvalidInteger,
		},
		"NaN": {
			amount: "ten",
			err:    ErrInvalidInteger,
		},
		"empty string": {
			amount: "",
			err:    ErrInvalidInteger,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseNumber(tc.amount)
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
				integerPart: 15,
				decimalPart: 2,
				precision:   1,
			},
			expected: "15.2",
		},
		"15.02": {
			n: Number{
				integerPart: 15,
				decimalPart: 2,
				precision:   2,
			},
			expected: "15.02",
		},
		"15.0200": {
			n: Number{
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
