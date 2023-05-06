package isomorphic

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegression(t *testing.T) {
	tests := []struct {
		input    []string
		expected bool
	}{
		{
			input:    []string{"egg", "add"},
			expected: true,
		},
		{
			input:    []string{"foo", "bar"},
			expected: false,
		},
		{
			input:    []string{"paper", "title"},
			expected: true,
		},
		{
			input:    []string{"a", "b"},
			expected: true,
		},
		{
			input:    []string{"a", "a"},
			expected: true,
		},
		{
			input:    []string{"badc", "baba"},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(strings.Join(test.input, " "), func(t *testing.T) {
			assert.Equal(t, test.expected, isIsomorphic(test.input[0], test.input[1]))
		})
	}
}
