package main

import "fmt"

func main() {
	total := sumAll(10, 30, 50, 80, 90)
	fmt.Println("Total penjumlahan =", total)

	// slice parameter
	slice := []int{10, 20, 30}
	total = sumAll(slice...)
	fmt.Println("Total penjumlahan menggunakan slice =", total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value

	}
	return total
}
