package main

import "fmt"

func main() {
	sayHelloTo("Sahril", "Mahendra")
}

func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}
