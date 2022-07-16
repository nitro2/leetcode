package palindrome

import (
	"fmt"
	"testing"
)

func TestRegression(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{input: 131, expected: true},
		{input: 10, expected: false},
		{input: 212, expected: true},
		{input: 123454321, expected: true},
		{input: 12344321, expected: true},
		{input: 1000021, expected: false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			if got := isPalindrome(test.input); got != test.expected {
				t.Errorf("expected %t, got %t", test.expected, got)
			}
		})
	}
}
