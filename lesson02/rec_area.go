package main

import (
	"fmt"
)

func main() {
	var (
		a, b int
	)
	fmt.Print("Введите длину стороны a: ")
	fmt.Scan(&a)

	fmt.Print("Введите длину стороны b: ")
	fmt.Scan(&b)

	fmt.Printf("Площадь прямоугольника со сторонами a = %d и b = %d составляет %d\n", a, b, a*b)
}
