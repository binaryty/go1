package main

import "fmt"

func simpleFibo(n int) uint64 {
	if n < 2 {
		return uint64(n)
	}
	return simpleFibo(n-1) + simpleFibo(n-2)
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d ", simpleFibo(i))
	}
	fmt.Print("\n")
}
