// Package calculator implements basic arithmetic operations.
package calculator

// Sum returns the sum of two int32 as an int32.
func Sum(firstNumber, secondNumber int32) int32 {
	return firstNumber + secondNumber
}

// Average returns the average between two float32 as a float32.
func Average(firstNumber, secondNumber float32) float32 {
	return (firstNumber + secondNumber) / 2
}
