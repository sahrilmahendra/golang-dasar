package main

import "fmt"

func main() {
	runApplication()
}
func logging() {
	fmt.Println("Function logging")
}

// defer function = tetap menjalankan function meskipun program error
func runApplication() {
	defer logging()
	fmt.Println("Aplikasi dijalankan")
}
