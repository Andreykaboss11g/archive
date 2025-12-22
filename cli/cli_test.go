package cli

import (
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "!hello"},
		{"World", "!world"},
		{"Test", "!test"},
	}

	for _, test := range tests {
		result := prepareData(test.input)
		if result != test.expected {
			t.Errorf("Encode(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}
