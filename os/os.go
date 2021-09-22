package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argumen :", args)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname", name)
	} else {
		fmt.Println("Error", err.Error())
	}
}