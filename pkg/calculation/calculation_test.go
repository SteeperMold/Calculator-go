package calculation

import (
	"errors"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		expected    float64
		expectedErr error
	}{
		// Valid expressions
		{
			name:       "simple addition",
			expression: "2 + 3",
			expected:   5,
		},
		{
			name:       "simple subtraction",
			expression: "10 - 7",
			expected:   3,
		},
		{
			name:       "simple multiplication",
			expression: "4 * 5",
			expected:   20,
		},
		{
			name:       "simple division",
			expression: "20 / 4",
			expected:   5,
		},
		{
			name:       "mixed operators with precedence",
			expression: "2 + 3 * 4",
			expected:   14, // 3*4 = 12, then 2+12 = 14
		},
		{
			name:       "parentheses override precedence",
			expression: "(2 + 3) * 4",
			expected:   20, // 2+3 = 5, then 5*4 = 20
		},
		{
			name:       "nested parentheses",
			expression: "((1 + 2) * (3 + 4)) / 2",
			expected:   10.5, // 1+2=3, 3+4=7, 3*7=21, then 21/2 = 10.5
		},
		{
			name:       "complex expression",
			expression: "3 + 4 * 2 / (1 - 5)",
			expected:   1, // 4*2=8, 8/(1-5) = -2, then 3 + (-2) = 1
		},
		{
			name:       "expression with spaces",
			expression: " 2 + 3 * ( 4 - 1 ) ",
			expected:   11, // 4-1=3, 3*3=9, then 2+9=11
		},

		// Invalid expressions
		{
			name:        "empty expression",
			expression:  "",
			expectedErr: ErrInvalidExpression,
		},
		{
			name:        "invalid character in expression",
			expression:  "2 + a",
			expectedErr: ErrInvalidExpression,
		},
		{
			name:        "mismatched parentheses",
			expression:  "(2 + 3",
			expectedErr: ErrInvalidExpression,
		},
		{
			name:        "extra operator",
			expression:  "2 + + 3",
			expectedErr: ErrInvalidExpression,
		},
		{
			name:        "division by zero",
			expression:  "5 / 0",
			expectedErr: ErrInvalidExpression, // Handle in `applyOperator` if needed
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(tt.expression)

			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("Expected error: %v, got: %v", tt.expectedErr, err)
			}

			if err == nil && result != tt.expected {
				t.Fatalf("Expected result: %v, got: %v", tt.expected, result)
			}
		})
	}
}
