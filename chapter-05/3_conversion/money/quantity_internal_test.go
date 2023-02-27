package money

import (
	"errors"
	"testing"
)

func TestParseQuantity(t *testing.T) {
	tt := map[string]struct {
		quantity string
		expected Quantity
		err      error
	}{
		"2 decimal digits": {
			quantity: "1.52",
			expected: Quantity{152, 2},
			err:      nil,
		},
		"no decimal digits": {
			quantity: "1",
			expected: Quantity{1, 0},
			err:      nil,
		},
		"suffix 0 as decimal digits": {
			quantity: "1.50",
			expected: Quantity{150, 2},
			err:      nil,
		},
		"prefix 0 as decimal digits": {
			quantity: "1.02",
			expected: Quantity{102, 2},
			err:      nil,
		},
		"invalid decimal part": {
			quantity: "65.pocket",
			err:      ErrInvalidValue,
		},
		"with apostrophes for readability": {
			quantity: "12'152.03",
			// expected: Quantity{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
			err: ErrInvalidValue,
		},
		"with underscores for readability": {
			quantity: "12_152.03",
			// expected: Quantity{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
			err: ErrInvalidValue,
		},
		"NaN": {
			quantity: "ten",
			err:      ErrInvalidValue,
		},
		"empty string": {
			quantity: "",
			err:      ErrInvalidValue,
		},
		"too large": {
			quantity: "1234567890123",
			err:      ErrTooLarge,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseQuantity(tc.quantity)
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
