package alg

import "fmt"

// ExampleInsSort demonstrates sorting by InsSort function.
func ExampleInsSort() {
	arr := []int{13, 5, 9, 12, 13, 21, 8, 55, 32, 99}
	fmt.Println(InsSort(arr))
	// Output: [5 8 9 12 13 13 21 32 55 99]
}

// ExampleBubbleSort demonstrates sorting by BubbleSort function.
func ExampleBubbleSort() {
	arr := []int{13, 5, 9, 12, 13, 21, 8, 55, 32, 99}
	fmt.Println(BubbleSort(arr))
	// Output: [5 8 9 12 13 13 21 32 55 99]
}

// ExampleQuickSort demonstrates sorting by QuickSort function.
func ExampleQuickSort() {
	arr := []int{13, 5, 9, 12, 13, 21, 8, 55, 32, 99}
	fmt.Println(QuickSort(arr))
	// Output: [5 8 9 12 13 13 21 32 55 99]
}
