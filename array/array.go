package main

import "fmt"

func main() {
	var nama [3]string
	nama[0] = "Moch."
	nama[1] = "Syahryil"
	nama[2] = "Mahendra"

	fmt.Println(nama)
	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])
	fmt.Println(len(nama))

	var kota = [5]string{"Sidoarjo", "Solo", "Surabaya"}
	fmt.Println(kota)
	fmt.Println(kota[3])
	fmt.Println(len(kota))

}
