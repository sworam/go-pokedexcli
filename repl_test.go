package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " This is a string",
			expected: []string{"this", "is", "a", "string"},
		},
		{
			input:    "ALL UPPERCASE",
			expected: []string{"all", "uppercase"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths of the slizes dont match actual: %d, expected: %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words '%s' and '%s' dont match!", word, expectedWord)
			}
		}
	}
}
