package calculator

import (
	"errors"
	"math"
	"testing"
)

func TestCalculateSupportedOperations(t *testing.T) {
	tests := []struct {
		name      string
		operation string
		a         float64
		b         float64
		want      float64
	}{
		{name: "addition", operation: "add", a: 10, b: 5, want: 15},
		{name: "subtraction", operation: "subtract", a: 10, b: 5, want: 5},
		{name: "multiplication", operation: "multiply", a: 10, b: 5, want: 50},
		{name: "division", operation: "divide", a: 10, b: 4, want: 2.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.operation, tt.a, tt.b)
			if err != nil {
				t.Fatalf("Calculate() returned unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateDivisionByZero(t *testing.T) {
	_, err := Calculate("divide", 10, 0)
	if !errors.Is(err, ErrDivisionByZero) {
		t.Fatalf("Calculate() error = %v, want %v", err, ErrDivisionByZero)
	}
}

func TestCalculateUnsupportedOperation(t *testing.T) {
	_, err := Calculate("modulo", 10, 3)
	if !errors.Is(err, ErrUnsupportedOperation) {
		t.Fatalf("Calculate() error = %v, want %v", err, ErrUnsupportedOperation)
	}
}

func TestCalculateRejectsNonFiniteOperand(t *testing.T) {
	_, err := Calculate("add", math.Inf(1), 1)
	if !errors.Is(err, ErrInvalidOperand) {
		t.Fatalf("Calculate() error = %v, want %v", err, ErrInvalidOperand)
	}
}

func TestCalculateRejectsNonFiniteResult(t *testing.T) {
	_, err := Calculate("multiply", math.MaxFloat64, 2)
	if !errors.Is(err, ErrInvalidResult) {
		t.Fatalf("Calculate() error = %v, want %v", err, ErrInvalidResult)
	}
}
