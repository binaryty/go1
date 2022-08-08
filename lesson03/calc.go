package main

import "fmt"

func main() {
	var (
		a, b float32
		op   string
	)

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("введите второе число: ")
	fmt.Scan(&b)

	fmt.Print("Введите оператор: ")
	fmt.Scan(&op)

	switch op {
	case "+":
		fmt.Printf("Сумма чисел равна: %.0f\n", a+b)
	case "-":
		fmt.Printf("Разность чисел равна: %.0f\n", a-b)
	case "*":
		fmt.Printf("Произзведение чисел равно: %.0f\n", a*b)
	case "/":
		if b > 0 {
			fmt.Printf("Частное числе равно: %.2f\n", a/b)
		} else {
			fmt.Println("Деление на 0 запрещено!")
		}
	default:
		fmt.Println("Неизвестный оператор")
	}
}
