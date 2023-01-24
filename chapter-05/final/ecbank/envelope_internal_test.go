package ecbank

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ablqk/tiny-go-projects/chapter-05/layered/money"
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

	var ecbMessage Envelope

	expectedMessage := Envelope{
		Rates: []CurrencyRate{
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

	err := xml.Unmarshal([]byte(message), &ecbMessage)
	if err != nil {
		t.Errorf("unable to marshal: %s", err.Error())
	}

	assert.Equal(t, expectedMessage, ecbMessage)
}

func TestChangeRate(t *testing.T) {
	tt := map[string]struct {
		envelope Envelope
		source   money.Currency
		target   money.Currency
		want     money.ExchangeRate
		wantErr  error
	}{
		"nominal": {
			envelope: Envelope{Rates: []CurrencyRate{{Currency: "USD", Rate: 1.5}}},
			source:   money.NewCurrency("EUR", 2, 1),
			target:   money.NewCurrency("USD", 2, 1.5),
			want:     money.ExchangeRate(1.5),
			wantErr:  nil,
		}}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := tc.envelope.changeRate(tc.source, tc.target)
			if tc.wantErr != err {
				t.Errorf("unable to marshal: %s", err.Error())
			}
			if got != tc.want {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}
