package main

import "fmt"

func main() {
	// operator boolean pada golang
	// &&, ||, !
	var (
		nilaiAkhir = 90
		absensi    = 80

		lulusNilaiAkhir = nilaiAkhir > 75
		lulusAbsensi    = absensi > 70

		lulus = lulusNilaiAkhir && lulusAbsensi
	)
	fmt.Println(lulus)
}
