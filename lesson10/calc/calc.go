// Package calc is designed to perform operations of sum, subtraction, multiplication, division, power and factorial of numbers
package calc

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

const (
	OpOne    = "Input a number        : "
	OpFirst  = "Input a first number  : "
	OpSecond = "Input a second number : "
	OpBase   = "Input a base          : "
	OpPower  = "Input a power         : "
)

var (
	ErrorDivideByAero   = errors.New("can't divide by zero")
	ErrorNegativeNumber = errors.New("value can't  be less than zero")
)

type Calculator interface {
	Init(v ...float64)
	Calculate() (float64, error)
}

type Add struct {
	N float64
	M float64
}

type Sub struct {
	N float64
	M float64
}

type Mul struct {
	N float64
	M float64
}

type Div struct {
	N float64
	M float64
}

type Fact struct {
	N float64
}

type Pow struct {
	B float64
	P float64
}

func (a Add) Init(v ...float64) {
	a.N = v[0]
	a.M = v[1]
}

func (a Add) Calculate() (float64, error) {

	return a.N + a.M, nil
}

func (s Sub) Init(v ...float64) {
	s.N = v[0]
	s.M = v[1]
}

func (s Sub) Calculate() (float64, error) {

	return s.N - s.M, nil
}

func (mu Mul) Init(v ...float64) {
	mu.N = v[0]
	mu.M = v[1]
}

func (mu Mul) Calculate() (float64, error) {

	return mu.N * mu.M, nil
}

func (d Div) Init(v ...float64) {
	d.N = v[0]
	d.M = v[1]
}

func (d Div) Calculate() (float64, error) {
	if d.M == 0 {
		return 0, ErrorDivideByAero
	}

	return d.N / d.M, nil
}

func (f Fact) Init(v ...float64) {
	f.N = v[0]
}
func (f Fact) Calculate() (float64, error) {
	if f.N < 0 {
		return 0, ErrorNegativeNumber
	}

	return float64(factorial(int(f.N))), nil
}

func (p Pow) Init(v ...float64) {
	p.P = v[0]
	p.B = v[1]
}
func (p Pow) Calculate() (float64, error) {

	return math.Pow(p.B, p.P), nil
}

func Calculate(o Calculator) (float64, error) {

	return o.Calculate()
}

// factorial calculates the value of the factorial of a natural number
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

// GetNumber function accepts a value with the text of the user input and returns the read and validated value
func GetNumber(s string) float64 {
	var a string
	fmt.Print(s)
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("can't read")
	}
	n, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println("value is not a number")
		return 0
	}

	return n
}
