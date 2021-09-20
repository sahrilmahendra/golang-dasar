package main

import "fmt"

func main() {
	runApp(true)
}
func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi dihentikan")
	}
	fmt.Println("Aplikasi berjalan")
}
