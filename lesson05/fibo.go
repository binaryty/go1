package main

import "fmt"

var (
	cache = make(map[int]uint64)
)

func simpleFibo(n int) uint64 {
	if n < 2 {
		return uint64(n)
	}
	return simpleFibo(n-1) + simpleFibo(n-2)
}

func cachedFibo(n int) uint64 {
	if n < 2 {
		return uint64(n)
	}

	if value, ok := cache[n]; ok {
		return value
	}

	fiboResilt := cachedFibo(n-1) + cachedFibo(n-2)
	cache[n] = fiboResilt

	return fiboResilt
}

func main() {
	for i := 0; i < 25; i++ {
		// fmt.Printf("%d ", simpleFibo(i))
		fmt.Printf("%d ", cachedFibo(i))
	}
	fmt.Print("\n")
}
