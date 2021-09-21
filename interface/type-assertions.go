package main

import "fmt"

func random() interface{} {
	return "OK"
}
func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)
	switch value := result.(type) {
	case string:
		fmt.Println(value, "is string")
	case int:
		fmt.Println(value, "is int")
	default:
		fmt.Println(value, "is unknown")
	}
}
