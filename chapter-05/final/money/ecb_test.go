package money

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
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
		Cube: EnvelopeCube{
			ParentCube: ParentCube{
				Time: "2022-11-25",
				Cubes: []Cube{
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
			},
		},
	}

	err := xml.Unmarshal([]byte(message), &ecbMessage)
	if err != nil {
		t.Errorf("unable to marshal: %s", err.Error())
	}

	assert.Equal(t, expectedMessage, ecbMessage)
}
