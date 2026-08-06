package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input 		string
		expected	[]string
	}{
		{
			input: "     hello      world",
			expected: []string{"hello", "world"},
		},{
			input: "",
			expected: []string{},
		},{
			input: " HELLO  woRld      ",
			expected: []string{"hello", "world"},
		},{
			input: " hello World ",
			expected: []string{"hello", "world"},

		},{
			input: "        ",
			expected: []string{},
		},
	}
	
	for _, c := range cases {
		actual := cleanInput(c.input)
		
		if len(actual) != len(c.expected) {
			t.Errorf("mismatched arrays, expected length of %d, got %d", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if (word != expectedWord) {
				t.Errorf("mismatched words, expected '%s', got '%s'", expectedWord, word)
			}
		}
	}


}
