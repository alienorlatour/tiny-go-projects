package main

import "testing"

func Test_input(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := input(); got != tt.want {
				t.Errorf("input() = %v, want %v", got, tt.want)
			}
		})
	}
}
