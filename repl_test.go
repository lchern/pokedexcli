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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Go is Awesome!",
			expected: []string{"go", "is", "awesome!"},
		},
		{
			input:    "ALL CAPS AND SYMBOLS!!!@#$",
			expected: []string{"all", "caps", "and", "symbols!!!@#$"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "New Line\nTab\tSpace ",
			expected: []string{"new", "line", "tab", "space"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected length %d but got %d", c.input, len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Fatalf("For input '%s', expected word '%s' but got '%s'", c.input, expectedWord, word)
			}
		}
	}
}