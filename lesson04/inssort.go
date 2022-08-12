package main

import "fmt"

func main() {
	var arr = []int{2, 7, 4, 0, 9, 5, 3, 8, 6, 1}

	fmt.Printf("%v\n", insertionSort(arr))
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		el := arr[i]
		idx := i
		for ; idx > 0 && arr[idx-1] > el; idx-- {
			arr[idx] = arr[idx-1]
		}
		arr[idx] = el
	}
	return arr
}
