package main

import (
	"fmt"
	"github.com/binaryty/go1/lesson10/alg"
	"github.com/binaryty/go1/lesson10/area"
)

func main() {

	var arr = make([]int, 10)

	alg.FillArr(arr)
	fmt.Printf("\nUnsorted data: %v\n", arr)
	fmt.Printf("  Sorted data: %v\n\n", alg.InsSort(arr))

	square := area.Square{SideLength: 5}
	rect := area.Rectangle{SideA: 4, SideB: 7}
	circle := area.Circle{Radius: 5}
	circleNegative := area.Circle{Radius: -5}
	shapes := []area.Shape{square, rect, circle, circleNegative}

	for _, s := range shapes {
		area.PrintArea(s)
	}
}
