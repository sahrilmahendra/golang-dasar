package main

import "fmt"

func main() {
	nama := "mahendra"

	switch nama {
	case "sahril":
		fmt.Println("Hello sahril")
	case "mahendra":
		fmt.Println("Hello mahendra")
	default:
		fmt.Println("Siapa namamu..?")
	}
}
