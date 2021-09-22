package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hasil strings.Contains :", strings.Contains("Sahril Mahendra", "Sahril"))
	fmt.Println("Hasil strings.Contains :", strings.Contains("Sahril Mahendra", "Yadi"))

	fmt.Println("Hasil strings.Split :", strings.Split("Moch. Syahryil Mahendra", " "))

	fmt.Println("Hasil strings.ToLower :", strings.ToLower("Moch. Syahryil Mahendra"))
	fmt.Println("Hasil strings.ToUpper :", strings.ToUpper("Moch. Syahryil Mahendra"))

	fmt.Println("Hasil strings.Trim :", strings.Trim("                        Sahril          ", " "))

	fmt.Println("Hasil strings.ReplaceAll :", strings.ReplaceAll("sahril sahril sahril", "sahril", "Mahendra"))
}
