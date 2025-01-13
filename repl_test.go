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
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Hello    World   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Public Private GREG",
			expected: []string{"public", "private", "greg"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of result is not correct.\n Expected %v\n Actual %v",
				len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Resultant slice is not correct.\n Expected %v\n Actual %v",
					expectedWord, word)
			}
		}
	}

}
