package main

import "fmt"

func main() {
	fmt.Println(fullName())
	firstName, middleName, lastName := fullName()
	fmt.Println("Hello", firstName, middleName, lastName)
}

func fullName() (firstName, middleName, lastName string) {
	firstName = "Moch."
	middleName = "Syahryil"
	lastName = "Mahendra"
	return
}
