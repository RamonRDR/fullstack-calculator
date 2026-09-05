package calculator

import (
	"errors"
	"math"
)

var (
	ErrDivisionByZero      = errors.New("division by zero")
	ErrInvalidOperand      = errors.New("operands must be finite numbers")
	ErrInvalidResult       = errors.New("result is not a finite number")
	ErrUnsupportedOperation = errors.New("unsupported operation")
)

// Calculate applies one supported binary operation to a and b.
func Calculate(operation string, a, b float64) (float64, error) {
	if !isFinite(a) || !isFinite(b) {
		return 0, ErrInvalidOperand
	}

	var result float64

	switch operation {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	case "divide":
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		result = a / b
	default:
		return 0, ErrUnsupportedOperation
	}

	if !isFinite(result) {
		return 0, ErrInvalidResult
	}

	return result, nil
}

func isFinite(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
}
