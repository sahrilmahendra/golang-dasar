package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 32768
		nilai64 int64 = int64(nilai32)

		nilai16 int16 = int16(nilai32)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var (
		name           = "Sahril"
		e       byte   = name[0]
		eString string = string(e)
	)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
