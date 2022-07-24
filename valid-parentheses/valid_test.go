package validParentheses

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input    string
	expected bool
}{
	{input: "()", expected: true},
	{input: "()]", expected: false},
	{input: "{[()[]]}", expected: true},
	{input: "{}([{}])", expected: true},
	{input: "{}([{}])[", expected: false},
	{input: ")", expected: false},
}

func TestValidParentheses(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%v expects %v", test.input, test.expected), func(t *testing.T) {
			got := isValid(test.input)
			if got != test.expected {
				t.Errorf("got %v, want %v", got, test.expected)
			}
		})
	}
}
