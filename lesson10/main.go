package main

import (
	"fmt"
	"github.com/binaryty/go1/lesson10/alg"
	"github.com/binaryty/go1/lesson10/area"
	"github.com/binaryty/go1/lesson10/calc"
)

func main() {
	fmt.Println("------------------alg pkg----------------")

	var arr = make([]int, 10)

	alg.FillArr(arr)
	fmt.Printf("Unsorted data: %v\n", arr)
	fmt.Printf("  Sorted data: %v\n", alg.InsSort(arr))

	fmt.Println("-----------------area pkg-----------------")

	square := area.Square{SideLength: 5}
	rect := area.Rectangle{SideA: 4, SideB: 7}
	circle := area.Circle{Radius: 5}
	circleNegative := area.Circle{Radius: -5}
	shapes := []area.Shape{square, rect, circle, circleNegative}
	for _, s := range shapes {
		area.PrintArea(s)
	}

	fmt.Println("----------------calc pkg------------------")

	var op string
	var o calc.Calculator

	fmt.Printf("Input an operation    : ")
	_, err := fmt.Scan(&op)
	if err != nil {
		fmt.Printf("Can't scan operation: %v", err)
	}

	switch op {
	case "+":
		o = calc.Add{N: calc.GetNumber(calc.OpFirst), M: calc.GetNumber(calc.OpSecond)}
	case "-":
		o = calc.Sub{N: calc.GetNumber(calc.OpFirst), M: calc.GetNumber(calc.OpSecond)}
	case "*":
		o = calc.Mul{N: calc.GetNumber(calc.OpFirst), M: calc.GetNumber(calc.OpSecond)}
	case "/":
		o = calc.Div{N: calc.GetNumber(calc.OpFirst), M: calc.GetNumber(calc.OpSecond)}
	case "^":
		o = calc.Pow{B: calc.GetNumber(calc.OpBase), P: calc.GetNumber(calc.OpPower)}
	case "!":
		o = calc.Fact{N: calc.GetNumber(calc.OpOne)}
	default:
		fmt.Println("Wrong operator")
		return
	}
	res, err := o.Calculate()
	if err != nil {
		fmt.Printf("Error in %T func: %v\n", o, err)
	} else {
		fmt.Printf("Result                : %v\n", res)
	}
}
