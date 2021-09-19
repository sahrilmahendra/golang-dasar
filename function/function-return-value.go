package main

import "fmt"

func main() {
	fmt.Println(getHello("Sahril Mahendra"))
}
func getHello(nama string) string {
	return "Hello " + nama
}
