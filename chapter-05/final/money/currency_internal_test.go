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

	got, err := ParseCurrency("EUR")
	if err != nil {
		t.Errorf("expected no error, got %s", err.Error())
	}

	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestGetCurrency_InvalidCurrencyCode(t *testing.T) {
	_, err := ParseCurrency("INVALID")
	if !errors.Is(err, errInvalidCurrencyCode) {
		t.Errorf("expected error %s, got %v", errInvalidCurrencyCode, err)
	}
}
