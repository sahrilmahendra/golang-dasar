package main

import "fmt"

func main() {
	// nama := "Sahril"
	// if nama == "Sahril" {
	// 	fmt.Println("Hello,", nama)
	// } else if nama == "Mahendra" {
	// 	fmt.Println("Hello", nama)
	// } else {
	// 	fmt.Println("Hello, Siapa namamu..?")
	// }
	if nama := "mahendra"; nama == "sahril" {
		fmt.Println("Hello,", nama)
	} else if nama == "mahendra" {
		fmt.Println("apakah namamu", nama, "..?")
	} else {
		fmt.Println("Hello, Siapa namamu..?")
	}

}
