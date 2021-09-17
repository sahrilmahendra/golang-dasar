package main

import "fmt"

func main() {
	// operator matematika dalam golang +, -, *, /, %
	var (
		a     = 1
		b     = 2
		hasil = a + b
	)
	fmt.Println("a sama dengan ", a)
	fmt.Println("b sama dengan ", b)
	fmt.Println("a + b hasilnya ", hasil)

	// augmented assignment
	// a += 10 || a = a + 10
	// a -= 10 || a = a - 10
	// a *= 10 || a = a * 10
	// a /= 10 || a = a / 10
	// a %= 10 || a = a % 10

	var x = 10
	x += 10 // x = x + 10
	fmt.Println("x = 10, maka x += 10 hasilnya", x)

	// unary operator
	// a++ || a = a + 1
	// a-- || a = a - 1
	// negative (-)
	// positive (+)
	// boolean negasi (!)
	var i = 2
	i++
	fmt.Println("i = 2, maka i++ hasilnya", i)
}
