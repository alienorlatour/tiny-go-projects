package money

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewAmount(t *testing.T) {
	tt := map[string]struct {
		quantity Decimal
		currency Currency
		want     Amount
		err      error
	}{
		"1.50 €": {
			quantity: Decimal{subunits: 150, precision: 2},
			currency: Currency{code: "EUR", precision: 2},
			want: Amount{
				quantity: Decimal{subunits: 150, precision: 2},
				currency: Currency{code: "EUR", precision: 2},
			},
		},
		"1.500 €": {
			quantity: Decimal{subunits: 1500, precision: 3},
			currency: Currency{code: "EUR", precision: 2},
			err:      ErrTooPrecise,
		},
		"1.5 €": {
			quantity: Decimal{subunits: 15, precision: 1},
			currency: Currency{code: "EUR", precision: 2},
			want: Amount{
				quantity: Decimal{subunits: 150, precision: 2},
				currency: Currency{code: "EUR", precision: 2},
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := NewAmount(tc.quantity, tc.currency)
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}
