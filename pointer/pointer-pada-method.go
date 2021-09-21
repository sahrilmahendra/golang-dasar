package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	rudi := Man{"Rudi"}
	rudi.Married()
	fmt.Println(rudi.Name)
}
