package factorial

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	// Test cases table
	testCases := []struct {
		name     string
		input    int
		expected uint64
		err      error
	}{
		{"factorial of 0", 0, 1, nil},
		{"factorial of 1", 1, 1, nil},
		{"factorial of 5", 5, 120, nil},
		{"error case: negative input", -1, 0, ErrNegative},
		{"error case: overflow", 84, 0, ErrOverflow},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { // Run each test case in a separate subtest
			result, err := Factorial(tc.input)
			if err != tc.err {
				t.Errorf("For input %d, expected error: %v, got error: %v", tc.input, tc.err, err)
			}
			if result != tc.expected {
				t.Errorf("For input %d, expected: %d, got: %d", tc.input, tc.expected, result)
			}
		})
	}
}
