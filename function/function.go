package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		sayHello()
	}
}

func sayHello() {
	fmt.Println("Hello World")
}
