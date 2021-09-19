package main

import "fmt"

func main() {
	hello := hello
	hasil := hello("Sahril")
	fmt.Println(hasil)
}

func hello(nama string) string {
	return "Hello " + nama
}
