package tdd

import "errors"

// Add returns sum of two integers
func Add(a, b int) int {
	return a + b
}

// Subtract returns difference of two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide returns quotient of two integers, errors if divisor is zero
func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return float64(a) / float64(b), nil
}
