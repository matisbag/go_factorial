package factorial

import "errors"

// Define custom error types for specific error scenarios
var (
	ErrNegative = errors.New("negative input not allowed")
	ErrOverflow = errors.New("factorial overflow")
)

func Factorial(n int) (uint64, error) {
	// If the input is negative, return an error
	if n < 0 {
		return 0, ErrNegative
	}
	// The factorial of 0 is 1
	if n == 0 {
		return 1, nil
	}

	// Call the recursive helper function to calculate the factorial
	result, err := factorialRecursive(uint64(n))
	if err != nil {
		return 0, err
	}

	return result, nil
}

func factorialRecursive(n uint64) (uint64, error) {
	// The factorial of 1 and 0 is 1
	if n <= 1 {
		return 1, nil
	}

	// Recursive call to calculate the factorial
	result, err := factorialRecursive(n - 1)
	if err != nil {
		return 0, err
	}

	// Check for overflow
	if n > uint64(^uint64(0))/result {
		return 0, ErrOverflow
	}

	// Return the calculated factorial
	return n * result, nil
}
