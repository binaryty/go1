package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var (
		op1, op2, operation     string
	)

	repeat := true
	for repeat {
		fmt.Print("Введите первое число: ")
		fmt.Scan(&op1)
		fA, err := strconv.ParseFloat(op1, 64)
		if err != nil {
			fmt.Println("Введено не число,  повторите попытку")
			continue
		}

		fmt.Print("Введите второе число: ")
		fmt.Scan(&op2)
		fB, err := strconv.ParseFloat(op2, 64)
		if err != nil {
			fmt.Println("Введено не число,  повторите попытку")
			continue
		}

		fmt.Print("Выберите оператор из списка(+-*/^!): ")
		fmt.Scan(&operation)

		switch operation {
		case "+":
			fmt.Printf("%.2f + %.2f = %.2f\n", fA, fB, fA+fB)
		case "-":
			fmt.Printf("%.2f  */- %.2f = %.2f\n", fA, fB, fA-fB)
		case "*":
			fmt.Printf("%.2f * %.2f = %.2f\n", fA, fB, fA*fB)
		case "/":
			if fB == 0 {
				fmt.Println("Деление на 0 запрещено")
				break
			} else {
				fmt.Printf("%.2f / %.2f = %.2f\n", fA, fB, fA/fB)
			}
		case "^":
			fmt.Printf("%.2f ^ %.2f = %.2f\n", fA, fB, math.Pow(fA, fB))
		case "!":
			iA, err := strconv.ParseUint(op1, 10, 64)
			iB, err := strconv.ParseUint(op2, 10, 64)
			if err != nil {
				fmt.Println("Факториал определен только для целых неотрицательных чисел")
				continue
			}
			if fA < 0 || fB < 0 {
				fmt.Println("Факториал определен только для целых неотрицательных чисел")
				continue
			} else {
			}
			fmt.Printf("!%d = %d и !%d = %d\n", iA, fac(iA), iB, fac(iB))
		default:
			repeat = false
			fmt.Println("Операция не найдена.")
		}
	}
}

func  fac(n uint64) uint64 {
	if n == 0 {
		return 1
	} else {
		return n * fac(n - 1)
	}
}
