package main

import "fmt"

func main() {
	runApp(false)
}

func endApp() {
	fmt.Println("Aplikasi selesai")
	// recover = menangkap nilai panic jika function panic terjadi
	pesan := recover()
	if pesan != nil {
		fmt.Println("Error :", pesan)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		// panic = fungsi yang menghentikan program jika terjadi error
		panic("Aplikasi dihentikan")
	}
	fmt.Println("Aplikasi berjalan")
}
