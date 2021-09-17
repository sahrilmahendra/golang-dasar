package main

import "fmt"

func main() {
	// operator perbandingan pada golang
	// > , >=, <, <=, ==, !=
	var (
		nama1 = "sahril"
		nama2 = "mahendra"
		hasil = nama1 == nama2
	)

	fmt.Println("nama1 == nama2, hasilnya", hasil)
}
