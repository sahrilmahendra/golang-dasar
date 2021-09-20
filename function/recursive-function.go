package main

import "fmt"

func main() {
	fmt.Println(factorialLoop(8))
	fmt.Println(factorialRecursive(5))
}

// faktorial menggunakan loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		fmt.Print(i)
		if i != 1 {
			fmt.Print(" * ")
		} else {
			fmt.Print(" = ")
		}
		result *= i
	}
	return result
}

// faktorial menggunakan recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
