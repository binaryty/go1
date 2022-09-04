package main

import (
	"fmt"
	"github.com/binaryty/go1/lesson10/alg"
)

func main() {

	var arr = make([]int, 10)

	alg.FillArr(arr)
	fmt.Printf("\nUnsorted data: %v\n", arr)
	fmt.Printf("  Sorted data: %v\n", alg.InsSort(arr))
}
