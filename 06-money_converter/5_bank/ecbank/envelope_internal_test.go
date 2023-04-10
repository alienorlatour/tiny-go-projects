package ecbank

import (
	"encoding/xml"
	"reflect"
	"testing"

	"learngo-pockets/moneyconverter/money"
)

func TestUnmarshalXML(t *testing.T) {
	message := `<gesmes:Envelope xmlns:gesmes="http://www.gesmes.org/xml/2002-08-01" xmlns="http://www.ecb.int/vocabulary/2002-08-01/eurofxref">
	<gesmes:subject>Reference rates</gesmes:subject>
	<gesmes:Sender>
		<gesmes:name>European Central Bank</gesmes:name>
	</gesmes:Sender>
	<Cube>
		<Cube time="2022-11-25">
			<Cube currency="JPY" rate="144.62"/>
			<Cube currency="TRY" rate="19.3333"/>
			<Cube currency="KRW" rate="1383.20"/>
			<Cube currency="NZD" rate="1.6651"/>
		</Cube>
	</Cube>
</gesmes:Envelope>`

	var got envelope

	want := envelope{
		Rates: []currencyRate{
			{
				Currency: "JPY",
				Rate:     144.62,
			},
			{
				Currency: "TRY",
				Rate:     19.3333,
			},
			{
				Currency: "KRW",
				Rate:     1383.20,
			},
			{
				Currency: "NZD",
				Rate:     1.6651,
			},
		},
	}

	err := xml.Unmarshal([]byte(message), &got)
	if err != nil {
		t.Errorf("unable to marshal: %s", err.Error())
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestChangeRate(t *testing.T) {
	tt := map[string]struct {
		envelope envelope
		source   string
		target   string
		want     money.ExchangeRate
		wantErr  error
	}{
		"nominal": {
			envelope: envelope{Rates: []currencyRate{{Currency: "USD", Rate: 1.5}}},
			source:   "EUR",
			target:   "USD",
			want:     mustParseExchangeRate(t, "1.5"),
			wantErr:  nil,
		},
		// TODO this is not enough
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := tc.envelope.exchangeRate(tc.source, tc.target)
			if tc.wantErr != err {
				t.Errorf("unable to marshal: %s", err.Error())
			}
			if got != tc.want {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}

func mustParseExchangeRate(t *testing.T, rate string) money.ExchangeRate {
	t.Helper()

	exchRate, err := money.ParseDecimal(rate)
	if err != nil {
		t.Fatalf("unable to parse exchange rate %s", rate)
	}
	return money.ExchangeRate(exchRate)
}
