package repository

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalXML(t *testing.T) {
	message := `<gesmes:envelope xmlns:gesmes="http://www.gesmes.org/xml/2002-08-01" xmlns="http://www.ecb.int/vocabulary/2002-08-01/eurofxref">
	<gesmes:subject>Reference rates</gesmes:subject>
	<gesmes:Sender>
		<gesmes:name>European Central Bank</gesmes:name>
	</gesmes:Sender>
	<cube>
		<cube time="2022-11-25">
			<cube currency="JPY" rate="144.62"/>
			<cube currency="TRY" rate="19.3333"/>
			<cube currency="KRW" rate="1383.20"/>
			<cube currency="NZD" rate="1.6651"/>
		</cube>
	</cube>
</gesmes:envelope>`

	var ecbMessage envelope

	expectedMessage := envelope{
		Cube: envelopeCube{
			ParentCube: parentCube{
				Time: "2022-11-25",
				Cubes: []cube{
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
