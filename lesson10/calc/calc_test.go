package calc_test

import (
	"errors"
	"math"
	"testing"

	"github.com/binaryty/go1/lesson10/calc"
)

func TestCalculate(t *testing.T) {
	TestCases := []struct {
		name     string
		input    calc.Calculator
		output   float64
		expError error
	}{
		{
			name:   "Add case with 0",
			input:  calc.Add{N: 0, M: 0},
			output: 0,
		},
		{
			name:   "Add case with natural numbers",
			input:  calc.Add{N: 15, M: 23},
			output: 38,
		},
		{
			name:   "Add case with two negative numbers",
			input:  calc.Add{N: -12, M: -45},
			output: -57,
		},
		{
			name:   "Add case with one negative number",
			input:  calc.Add{N: 12, M: -45},
			output: -33,
		},
		//-----------------------------------------------------
		{
			name:   "Sub case with 0",
			input:  calc.Sub{N: 0, M: 0},
			output: 0,
		},
		{
			name:   "Sub case with natural numbers",
			input:  calc.Sub{N: 23, M: 15},
			output: 8,
		},
		{
			name:   "Sub case with two negative numbers",
			input:  calc.Sub{N: -12, M: -45},
			output: 33,
		},
		{
			name:   "Sub case with one negative number",
			input:  calc.Sub{N: 12, M: -45},
			output: 57,
		},
		//-----------------------------------------------------
		{
			name:   "Mul case with 0",
			input:  calc.Mul{N: 0, M: 0},
			output: 0,
		},
		{
			name:   "Mul case with 0  and any number",
			input:  calc.Mul{N: 0, M: 5},
			output: 0,
		},
		{
			name:   "Mul case with 0 and negative number",
			input:  calc.Mul{N: 0, M: -5},
			output: 0,
		},
		{
			name:   "Mul case with natural numbers",
			input:  calc.Mul{N: 23, M: 15},
			output: 345,
		},
		{
			name:   "Mul case with two negative numbers",
			input:  calc.Mul{N: -12, M: -45},
			output: 540,
		},
		{
			name:   "Mul case with one negative number",
			input:  calc.Mul{N: 12, M: -45},
			output: -540,
		},
		//-----------------------------------------------------
		{
			name:     "Div case with 0",
			input:    calc.Div{N: 0, M: 0},
			output:   0,
			expError: calc.ErrorDivideByAero,
		},
		{
			name:   "Div case with 0  and any number",
			input:  calc.Div{N: 0, M: 5},
			output: 0,
		},
		{
			name:   "Div case with 0 and negative number",
			input:  calc.Div{N: 0, M: -5},
			output: 0,
		},
		{
			name:   "Div case with natural numbers",
			input:  calc.Div{N: 23, M: 15},
			output: 1.5333333333333334,
		},
		{
			name:   "Div case with two negative numbers",
			input:  calc.Div{N: -45, M: -15},
			output: 3,
		},
		{
			name:   "Div case with one negative number",
			input:  calc.Div{N: 45, M: -15},
			output: -3,
		},
		//-----------------------------------------------------
		{
			name:   "Pow case with 0",
			input:  calc.Pow{B: 0, P: 0},
			output: 1,
		},
		{
			name:   "Pow case with 0  and any number",
			input:  calc.Pow{B: 0, P: 5},
			output: 0,
		},
		{
			name:   "Pow case with 0 and negative number",
			input:  calc.Pow{B: 0, P: -5},
			output: math.Inf(1),
		},
		{
			name:   "Pow case with natural numbers",
			input:  calc.Pow{B: 2, P: 10},
			output: 1024,
		},
		{
			name:   "Pow case with two negative numbers",
			input:  calc.Pow{B: -2, P: -10},
			output: 0.0009765625,
		},
		{
			name:   "Pow case with one negative number",
			input:  calc.Pow{B: 10, P: -2},
			output: 0.01,
		},
		{
			name:   "Pow case with real numbers",
			input:  calc.Pow{B: 1.2, P: 2.5},
			output: 1.5774409656148782,
		},
		//-----------------------------------------------------
		{
			name:   "Fact case with 0",
			input:  calc.Fact{N: 0},
			output: 1,
		},
		{
			name:   "Fact case with 1",
			input:  calc.Fact{N: 1},
			output: 1,
		},
		{
			name:     "Fact case with negative number",
			input:    calc.Fact{N: -5},
			output:   0,
			expError: calc.ErrorNegativeNumber,
		},
		{
			name:   "Fact case with big number",
			input:  calc.Fact{N: 10},
			output: 3628800,
		},
	}

	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := calc.Calculate(tc.input)
			if !errors.Is(err, tc.expError) {
				t.Errorf("function must be return an error")
			}
			if result != tc.output {
				t.Errorf("expected: %v, got: %v\n", tc.output, result)
			}
		})
	}
}
