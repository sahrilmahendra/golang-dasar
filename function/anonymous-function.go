package main

import "fmt"

func main() {
	// anonymous function
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("admin", blacklist)
	registerUser("anjing", blacklist)
	registerUser("sahril", blacklist)
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println(name, "You are blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}
