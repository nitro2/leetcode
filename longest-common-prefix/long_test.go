package long

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegression(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Empty 1 input",
			input:    []string{""},
			expected: "",
		},
		{
			name:     "1 input",
			input:    []string{"a"},
			expected: "a",
		},
		{
			name:     "2 input",
			input:    []string{"ab", "ab"},
			expected: "ab",
		},
		{
			name:     "3 input",
			input:    []string{"ab", "ab", "abc"},
			expected: "ab",
		},
		{
			name:     "3 input 2",
			input:    []string{"abcd", "ab", "abc"},
			expected: "ab",
		},
		{
			name:     "First input empty",
			input:    []string{"", "ab", "abc"},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, longestCommonPrefix(test.input))
		})
	}
}
