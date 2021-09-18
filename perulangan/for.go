package main

import (
	"fmt"
)

func main() {
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("Perulangan ke", i)
	// 	i++
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Hello World", i)
	// }

	// sliceKota := []string{"Sidoarjo", "Surabaya", "Malang", "Mojokerto"}
	// // menggunakan for
	// // for i := 0; i <= len(sliceKota); i++ {
	// // 	fmt.Println(sliceKota)
	// // }
	// // menggunakan for range
	// for _, kota := range sliceKota {
	// 	fmt.Println(_, kota)
	// }
	mhs := make(map[string]string)
	mhs["nama"] = "Sahril Mahendra"
	mhs["nim"] = "123456"
	mhs["alamat"] = "Sidoarjo"
	for key, value := range mhs {
		fmt.Println(key, value)
	}
}
