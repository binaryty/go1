package area

import (
	"errors"
	"fmt"
	"math"
)

var ErrorNegativeValue = errors.New("value must be more than zero")

type Shape interface {
	Area() (float64, error)
}

type Square struct {
	SideLength float64
}

type Rectangle struct {
	SideA float64
	SideB float64
}

type Circle struct {
	Radius float64
}

func (s Square) Area() (float64, error) {
	if s.SideLength > 0 {
		return s.SideLength * s.SideLength, nil
	}
	return 0, ErrorNegativeValue
}

func (c Circle) Area() (float64, error) {
	if c.Radius > 0 {
		return math.Pi * math.Pow(c.Radius, 2), nil
	}
	return 0, ErrorNegativeValue
}

func (r Rectangle) Area() (float64, error) {
	if r.SideA > 0 && r.SideB > 0 {
		return r.SideA * r.SideB, nil
	}
	return 0, ErrorNegativeValue
}

func Area(s Shape) (float64, error) {
	return s.Area()
}

func PrintArea(s Shape) {
	switch s.(type) {
	case Square:
		res, err := s.Area()
		if err != nil {
			fmt.Printf("Error calculate %T: %v\n", s, err)
			return
		}
		fmt.Printf("   Square area = %f\n", res)
	case Rectangle:
		res, err := s.Area()
		if err != nil {
			fmt.Printf("Error calculate %T: %v\n", s, err)
			return
		}
		fmt.Printf("Rectangle area = %f\n", res)
	case Circle:
		res, err := s.Area()
		if err != nil {
			fmt.Printf("Error calculate %T: %v\n", s, err)
			return
		}
		fmt.Printf("   Circle area = %f\n", res)
	default:
		fmt.Printf("Unknown shape")
		return
	}
}
