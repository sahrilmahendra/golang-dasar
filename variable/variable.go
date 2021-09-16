package main

import "fmt"

func main(){
	// deklarasi variable secara utuh
	var name string

	name = "Moch. Syahryil"
	fmt.Println(name)

	name = "Sahril Mahendra"
	fmt.Println(name)

	// deklarasi variable tanpa menyebutkan tipe data
	// asalkan variable tsb langsung diberikan nilai
	var address = "Sidoarjo"
	fmt.Println(address)

	var age = 23
	fmt.Println(age)

	// deklarasi tanpa menggunakan var
	country := "Indonesia"
	fmt.Println(country)

	// deklarasi menggunakan multiple variable
	var(
		firstName = "Sahril"
		lastName = "Mahendra"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}