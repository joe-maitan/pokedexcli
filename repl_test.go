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
	} // End cases struct

	for _, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected

		if len(actual) != len(expected) {
			t.Errorf("Expected len of slice to be %v, but got %v", expected, actual)
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected %v, but got %v", expectedWord, word)
				t.Fail()
			}
		}
	} // End for loop
} // End TestCleanInput(testing.T)