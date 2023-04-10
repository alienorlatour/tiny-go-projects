package money

import (
	"errors"
	"testing"
)

func TestParseDecimal(t *testing.T) {
	tt := map[string]struct {
		decimal  string
		expected Decimal
		err      error
	}{
		"2 decimal digits": {
			decimal:  "1.52",
			expected: Decimal{152, 2},
			err:      nil,
		},
		"no decimal digits": {
			decimal:  "1",
			expected: Decimal{1, 0},
			err:      nil,
		},
		"suffix 0 as decimal digits": {
			decimal:  "1.50",
			expected: Decimal{15, 1},
			err:      nil,
		},
		"prefix 0 as decimal digits": {
			decimal:  "1.02",
			expected: Decimal{102, 2},
			err:      nil,
		},
		"multiple of 10": {
			decimal:  "150",
			expected: Decimal{150, 0},
			err:      nil,
		},
		"invalid decimal part": {
			decimal: "65.pocket",
			err:     ErrInvalidDecimal,
		},
		"with apostrophes for readability": {
			decimal: "12'152.03",
			// expected: Decimal{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
			err: ErrInvalidDecimal,
		},
		"with underscores for readability": {
			decimal: "12_152.03",
			// expected: Decimal{integerPart: 12152, decimalPart: 3, toUnit: 2}, // for future implementations
			err: ErrInvalidDecimal,
		},
		"Not a number": {
			decimal: "NaN",
			err:     ErrInvalidDecimal,
		},
		"empty string": {
			decimal: "",
			err:     ErrInvalidDecimal,
		},
		"too large": {
			decimal: "1234567890123",
			err:     ErrTooLarge,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseDecimal(tc.decimal)
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
