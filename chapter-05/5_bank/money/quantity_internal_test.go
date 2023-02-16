package money

import (
	"errors"
	"testing"
)

func TestParseQuantity(t *testing.T) {
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
			got, err := ParseQuantity(tc.amount)
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestQuantityString(t *testing.T) {
	tt := map[string]struct {
		quantity Quantity
		expected string
	}{
		"15.2": {
			quantity: Quantity{
				cents: 152,
				exp:   1,
			},
			expected: "15.2",
		},
		"0.0152": {
			quantity: Quantity{
				cents: 152,
				exp:   4,
			},
			expected: "0.0152",
		},
		"152": {
			quantity: Quantity{
				cents: 152,
				exp:   0,
			},
			expected: "152",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.quantity.String()
			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}
