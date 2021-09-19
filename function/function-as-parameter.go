package main

import "fmt"

func main() {
	sayHelloWithFilter("Sahril", cekFilter)
	sayHelloWithFilter("anjing", cekFilter)
}

// type declaration
type Filter func(string) string

func sayHelloWithFilter(nama string, filter Filter) {
	nameFiltered := filter(nama)
	fmt.Println("Hello", nameFiltered)
}

func cekFilter(nama string) string {
	if nama == "anjing" {
		return "****"
	} else {
		return nama
	}
}
