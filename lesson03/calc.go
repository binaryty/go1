package main

import (
	"fmt"
)

func main() {
	var (
		a, b   float64
		op     string
		repeat bool = true
	)

	for repeat {
		fmt.Print("Введите первое число: ")
		fmt.Scan(&a)

		fmt.Print("Введите второе число: ")
		fmt.Scan(&b)

		fmt.Print("Выберите оператор из списка(+-*/^!): ")
		//\n+ сложение\n- вычитание\n* умножение\n/ деление\n^ возведение в степень\n! факторил):\n ")
		fmt.Scan(&op)

		switch op {
		case "+":
			fmt.Printf("%.0f + %.0f = %.0f\n", a, b, a+b)
		case "-":
			fmt.Printf("%.0f - %.0f = %.0f\n", a, b, a-b)
		case "*":
			fmt.Printf("%.0f * %.0f = %.0f\n", a, b, a*b)
		case "/":
			if b == 0 {
				fmt.Println("Деление на 0 запрещено")
				break
			} else {
				fmt.Printf("%.0f / %.0f = %.2f\n", a, b, a/b)
			}
		case "^":
			pow := 1.0
			for i := 1.0; i <= b; i++ {
				pow *= a
			}
			fmt.Printf("%.0f ^ %.0f = %.0f\n", a, b, pow)
		case "!":
			fa := 1.0
			fb := 1.0
			if a < 0 || b < 0 {

			} else {
				for i := 1.0; i <= a; i++ {
					fa *= i
				}
				for i := 1.0; i <= b; i++ {
					fb *= i
				}
			}
			fmt.Printf("!%.0f = %.0f и !%.0f = %.0f\n", a, fa, b, fb)
		default:
			repeat = false
			fmt.Println("Операция не найдена.")
		}
	}
}
