package main

import "testing"

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
	}

	for _, test := range tests {
		actual := cleanInput(test.input)

		// Add a check for the length of the slice matches the expected length
		// if not use t.Errorf("error message") and fail the test
		if len(actual) != len(test.expected) {
			t.Errorf("input length %d did not match expected length %d", len(actual), len(test.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := test.expected[i]
			// check each word in the slice - if they don't match use t.Errorf("error message") and fail test
			if word != expectedWord {
				t.Errorf("For input %q, at index %d: expected %q, but got %q", test.input, i, test.expected[i], word)
			}
		}
	}
}
