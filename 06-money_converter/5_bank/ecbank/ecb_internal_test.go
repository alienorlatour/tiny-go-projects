package ecbank

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"learngo-pockets/moneyconverter/money"
)

func TestEuroCentralBank_FetchExchangeRate_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope><Cube><Cube>
			<Cube currency='USD' rate='2'/>
			<Cube currency='RON' rate='6'/>
		</Cube></Cube></gesmes:Envelope>`)
	}))
	defer ts.Close()

	ecb := Client{
		url: ts.URL,
	}

	got, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency(t, "RON"))
	want := mustParseDecimal(t, "3")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if money.Decimal(got) != want {
		t.Errorf("FetchExchangeRate() got = %v, want %v", money.Decimal(got), want)
	}
}

func mustParseCurrency(t *testing.T, code string) money.Currency {
	t.Helper()

	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("cannot parse currency %s code", code)
	}

	return currency
}

func mustParseDecimal(t *testing.T, decimal string) money.Decimal {
	t.Helper()

	dec, err := money.ParseDecimal(decimal)
	if err != nil {
		t.Fatalf("cannot parse decimal %s", decimal)
	}

	return dec
}
