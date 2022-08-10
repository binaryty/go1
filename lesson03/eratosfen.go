package main

import "fmt"

func main() {
	var (
		n int
	)

	fmt.Print("Введите целое число N: ")
	fmt.Scan(&n)

	simpleDimple := make([]int, n+1)
	for idx :=range simpleDimple{
		simpleDimple[idx] = idx
	}

	fmt.Printf("Простые числа от 0 до %d\n", n)
	for i := 2; i <= n; i++ {
		if simpleDimple[i] != 0 {
			fmt.Printf("%d ", simpleDimple[i])
			for j := i * i; j <= n; j += i {
				simpleDimple[j] = 0
			}
		}
	}
	fmt.Print("\n")
}
