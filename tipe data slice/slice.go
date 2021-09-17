package main

import "fmt"

func main() {
	hari := [7]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	slice1 := hari[0:5]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[3] = "Hari-diubah"
	fmt.Println(hari)

	// buat slice baru tanpa deklarasi array
	sliceBaru := make([]string, 3, 5)
	sliceBaru[0] = "saya"
	sliceBaru[1] = "kamu"
	sliceBaru[2] = "dia"

	fmt.Println(sliceBaru)
	fmt.Println("Panjang slice", len(sliceBaru))
	fmt.Println("Kapasitas slice", cap(sliceBaru))
}
