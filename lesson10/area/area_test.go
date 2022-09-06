package area_test

import (
	"errors"
	"github.com/binaryty/go1/lesson10/area"
	"testing"
)

func TestArea(t *testing.T) {
	TestCases := []struct {
		name          string
		input         area.Shape
		output        float64
		expectedError error
	}{
		{
			name: "case with Square(side=0)",
			input: area.Square{
				SideLength: 0,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
		{
			name: "case with Square(side=5)",
			input: area.Square{
				SideLength: 5,
			},
			output: 25,
		},
		{
			name: "case with Square(negative side=-5)",
			input: area.Square{
				SideLength: -5,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
		{
			name: "case with Rectangle(SideA=0, SideB=0)",
			input: area.Rectangle{
				SideA: 0,
				SideB: 0,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
		{
			name: "case with Rectangle(SideA=0, SideB=5)",
			input: area.Rectangle{
				SideA: 0,
				SideB: 5,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
		{
			name: "case with Rectangle(SideA=5, SideB=0)",
			input: area.Rectangle{
				SideA: 5,
				SideB: 0,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
		{
			name: "case with Rectangle(SideA=5, SideB=10)",
			input: area.Rectangle{
				SideA: 5,
				SideB: 10,
			},
			output: 50,
		},
		{
			name: "case with Rectangle(SideA=10, SideB=5)",
			input: area.Rectangle{
				SideA: 10,
				SideB: 5,
			},
			output: 50,
		},
		{
			name: "case with Rectangle(SideA=-5, SideB=-5)",
			input: area.Rectangle{
				SideA: -5,
				SideB: -5,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
		{
			name: "case with Rectangle(SideA=-5, SideB=0)",
			input: area.Rectangle{
				SideA: -5,
				SideB: 0,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
		{
			name: "case with Rectangle(SideA=-5, SideB=5)",
			input: area.Rectangle{
				SideA: -5,
				SideB: 5,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
		{
			name: "case with Rectangle(SideA=0, SideB=-5)",
			input: area.Rectangle{
				SideA: 0,
				SideB: -5,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
		{
			name: "case with Rectangle(SideA=5, SideB=-5)",
			input: area.Rectangle{
				SideA: 5.0,
				SideB: -5,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
		{
			name: "case with Circle(radius=0)",
			input: area.Circle{
				Radius: 0,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
		{
			name: "case with Circle(radius=5)",
			input: area.Circle{
				Radius: 5,
			},
			output: 78.53981633974483,
		},
		{
			name: "case with Circle(negative radius=-5)",
			input: area.Circle{
				Radius: -5,
			},
			output:        0,
			expectedError: area.ErrorNegativeValue,
		},
	}

	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := area.Area(tc.input)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("function must be return an error")
			}
			if result != tc.output {
				t.Errorf("expected: %v, got: %v", tc.output, result)
			}
		})
	}

}
