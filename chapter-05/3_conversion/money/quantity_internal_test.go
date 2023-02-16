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
		"too large": {
			amount: "1234567890123",
			err:    ErrTooLarge,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseNumber(tc.amount)
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
