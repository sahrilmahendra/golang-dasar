package main

import "fmt"

func main() {
	// deklarasi struct
	type Mhs struct {
		Nama, Alamat string
		Umur         int
		Lulus        bool
	}
	var sahril Mhs
	sahril.Nama = "Sahril Mahendra"
	sahril.Alamat = "Sidoarjo"
	sahril.Umur = 23
	sahril.Lulus = true

	fmt.Println(sahril)
	fmt.Println("nama saya", sahril.Nama)
	fmt.Println("alamat", sahril.Alamat)
	fmt.Println("umur saya", sahril.Umur)
	fmt.Println("status lulus", sahril.Lulus)

	// deklarasi struct dengan cara lain
	ahmad := Mhs{
		Nama:   "Ahmad Ridwan",
		Alamat: "Surabaya",
		Umur:   21,
		Lulus:  false,
	}
	fmt.Println(ahmad)
}
