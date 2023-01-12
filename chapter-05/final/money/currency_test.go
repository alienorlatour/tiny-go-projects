package money

import (
	"errors"
	"testing"
)

func TestGetCurrency_EUR(t *testing.T) {
	expected := Currency{
		code:      "EUR",
		precision: 2,
	}

	got, err := parseCurrency("EUR")
	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}

	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestGetCurrency_UnknownCurrency(t *testing.T) {
	_, err := parseCurrency("UNKNOWN")
	if !errors.Is(err, ErrUnknownCurrency) {
		t.Errorf("expected error %s, got %v", ErrUnknownCurrency, err)
	}
}
