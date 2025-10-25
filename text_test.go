package main

import (
	"fmt"
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
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "HELLO WORLD    ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Println("----------")
		fmt.Println(actual)
		fmt.Println(c.expected)
		if len(actual) != len(c.expected) {
			t.Errorf("expected slice not similar to actual slice.")
		}

		for i := range actual {
			if actual[i] == " " {
				t.Errorf("whitespace from start not removed.")
			}
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected word not similar to actual word.")
			}
		}
	}
}
