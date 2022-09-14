// Package alg is designed to sort integer arrays in various ways
package alg

import (
	"math/rand"
	"sort"
)

func FillArr(a []int) {
	for i := 0; i < len(a); i++ {
		a[i] = i * rand.Intn(1000)
	}
}

// InsSort function takes a slice of int as input and returns it sorted in ascending order.
func InsSort(arr []int) []int {
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

// BubbleSort function takes a slice of int as input and return it sorted in ascending order.
func BubbleSort(arr []int) []int {
	needSort := false
	for {
		needSort = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				needSort = true
			}
		}
		if !needSort {
			break
		}
	}

	return arr
}

// QuickSort function use std func sort.Ints
func QuickSort(arr []int) []int {
	sort.Ints(arr)

	return arr
}

// Less function compare two values and return true if the first value is less, otherwise false
func Less(a, b int) bool {
	return a < b
}

// IsSort function return true if array is sorted
func IsSort(arr []int) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if Less(arr[i], arr[i-1]) {
			return false
		}
	}

	return true
}
