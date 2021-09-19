package main

import "fmt"

func main() {
	fmt.Println(fullName())

	// menghiraukan variable menggunakan (_) "garis bawah"
	firstname, _ := fullName()
	fmt.Println(firstname)
}

func fullName() (string, string) {
	return "Sahril", "Mahendra"
}
