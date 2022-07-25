package gordle

import "testing"

func TestGordle_Play(t *testing.T) {
	type fields struct {
		reader *bufio.Reader
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gordle{
				reader: tt.fields.reader,
			}
			g.Play()
		})
	}
}
