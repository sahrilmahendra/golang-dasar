package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		counter := 10
		counter--
		fmt.Print("counter scope local ", counter)
		fmt.Println("")
	}
	increment()
	increment()
	fmt.Print("counter scope global ", counter)
}
