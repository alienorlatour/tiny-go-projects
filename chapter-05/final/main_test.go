package main

import (
	"strings"
	"testing"
)

func TestValidateInputs(t *testing.T) {
	tt := map[string]struct {
		from    string
		to      string
		argc    int
		wantErr string
	}{
		"nominal": {
			from:    "USD",
			to:      "EUR",
			argc:    1,
			wantErr: "",
		},
		"missing from": {
			to:      "EUR",
			argc:    1,
			wantErr: "missing input currency",
		},
		"missing to": {
			from:    "USD",
			argc:    1,
			wantErr: "missing output currency",
		},
		"missing amount": {
			from:    "USD",
			to:      "EUR",
			argc:    0,
			wantErr: "invalid number of arguments, expecting only the amount to convert",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			// call validateInputs method
			err := validateInputs(tc.from, tc.to, tc.argc)
			// verifies there is no error if it's not expected
			if tc.wantErr == "" {
				// TODO rework it
				if err != nil {
					t.Errorf("expected %v, want %v", tc.wantErr, err)
				}
			} else if !strings.Contains(err.Error(), tc.wantErr) {
				// verifies the error message contains the wanted error message
				t.Errorf("expected %v, want %v", tc.wantErr, err)
			}
		})
	}

}
