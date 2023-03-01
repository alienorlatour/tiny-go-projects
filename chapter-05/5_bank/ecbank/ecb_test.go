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
			<Cube currency='RON' rate='4'/>
		</Cube></Cube></gesmes:Envelope>`)
	}))
	defer ts.Close()

	ecb := EuroCentralBank{
		path: ts.URL,
	}

	got, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency(t, "RON"))

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if got != 2 {
		t.Errorf("FetchExchangeRate() got = %v, want %v", got, 2)
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
