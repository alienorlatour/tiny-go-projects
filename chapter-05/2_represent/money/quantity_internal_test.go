package money

import (
	"errors"
	"testing"
)

func TestParseNumber(t *testing.T) {
	tt := map[string]struct {
		amount   string
		expected Quantity
		err      error
	}{
		"2 decimal digits": {
			amount:   "1.52",
			expected: Quantity{152, 2},
			err:      nil,
		},
		"no decimal digits": {
			amount:   "1",
			expected: Quantity{1, 0},
			err:      nil,
		},
		"suffix 0 as decimal digits": {
			amount:   "1.50",
			expected: Quantity{150, 2},
			err:      nil,
		},
		"prefix 0 as decimal digits": {
			amount:   "1.02",
			expected: Quantity{102, 2},
			err:      nil,
		},
		"invalid decimal part": {
			amount: "65.pocket",
			err:    ErrInvalidValue,
		},
		"with apostrophes for readability": {
			amount: "12'152.03",
			// expected: Quantity{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
			err: ErrInvalidValue,
		},
		"with underscores for readability": {
			amount: "12_152.03",
			// expected: Quantity{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
			err: ErrInvalidValue,
		},
		"NaN": {
			amount: "ten",
			err:    ErrInvalidValue,
		},
		"empty string": {
			amount: "",
			err:    ErrInvalidValue,
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
