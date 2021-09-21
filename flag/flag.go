package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Masukkan host")
	var user *string = flag.String("user", "root", "Masukkan user")
	var password *string = flag.String("password", "", "Masukkan password")

	flag.Parse()
	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
}
