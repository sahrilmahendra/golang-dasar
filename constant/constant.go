package main

import "fmt"

func main() {
	// constant isinya tidak dapat diubah
	// berbeda dengan variable, constant tidak akan error meskipun tidak digunakan
	const (
		firstName = "Sahril"
		lastName  = "Mahendra"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(firstName, lastName)
}
