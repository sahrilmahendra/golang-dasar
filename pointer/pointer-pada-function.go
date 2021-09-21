package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	alamat := Address{
		City:     "Sidoarjo",
		Province: "Jawa Timur",
		Country:  "",
	}
	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
