package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai float32, pembagi float32) (float32, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}
func main() {
	hasil, err := Pembagian(100, 3)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
