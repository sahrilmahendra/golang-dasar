package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		"Sidoarjo",
		"Jawa Timur",
		"Indonesia",
	}
	// pointer & diikuti nama variabelnya
	address2 := &address1
	address3 := &address1

	address2.City = "Malang"

	*address2 = Address{
		"Bandung", "Jawa Barat", "Indonesia",
	}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// membuat pointer dari awal (dengan data kosong)
	address4 := new(Address)
	fmt.Println(address4)
}
