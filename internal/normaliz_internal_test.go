package internal

import (
	"testing"
)

func TestNormalizGolangCodeSnipped(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic multiple new lines",
			input:    "var i int\n\n\nvar j int",
			expected: "var i int\nvar j int",
		},
		{
			name:     "basic multiple spaces",
			input:    " var   i  int ",
			expected: "var i int",
		},
		{
			name:     "basic multiple tabs",
			input:    "\t\tvar i int\t",
			expected: "var i int",
		},
		{
			name:     "basic spaces inside string",
			input:    "var k int\ns := \"Hello World\"",
			expected: "var k int\ns := \"Hello World\"",
		},
		{
			name:     "complex spaces inside string",
			input:    "var k int\ns := \"Hello \n  \\\" \t \n World\"",
			expected: "var k int\ns := \"Hello \n  \\\" \t \n World\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizGolangCodeSnipped(tt.input)
			if err != nil {
				t.Fatalf("NormalizeGoCode returned an error: %v", err)
			}
			if got != tt.expected {
				t.Errorf("NormalizeGoCode() = %q, want %q", got, tt.expected)
			}
		})
	}
}
