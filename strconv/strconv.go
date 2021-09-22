package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	number, err := strconv.ParseInt("0891201", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Error", err.Error())
	}

	hasil := strconv.FormatInt(931084, 2)
	fmt.Println(hasil)

	hasilInt, _ := strconv.Atoi("900123")
	fmt.Println(hasilInt)
}
