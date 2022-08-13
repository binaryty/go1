package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	arr := make([]int, 0, 100)

	fmt.Println("Введите элементы массива (целые числа через Enter) или q чтобы завершить ввод")

	for scanner.Scan() {
		input := scanner.Text()
		if input == "q" {
			break
		}
		num, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("Вы ввели не число! Повторите ввод.")
			continue
		}
		arr = append(arr, int(num))
	}
	fmt.Printf("Исходный массив:\n%v\n", arr)
	fmt.Printf("Отсортированный массив:\n%v\n", insertionSort(arr))
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
