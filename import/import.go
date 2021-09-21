package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	helper.SayHello("Sahril")
	// helper.sayGoodBye("Sahril") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
