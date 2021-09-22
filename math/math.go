package main

import (
	"fmt"
	"math"
)

func main() {
	// math.round
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.4))

	// math.Ceil
	fmt.Println(math.Ceil(1.1))
	fmt.Println(math.Ceil(1.7))

	// math.Floor
	fmt.Println(math.Floor(1.3))
	fmt.Println(math.Floor(1.6))

	// math.Max
	fmt.Println(math.Max(10, 20))

	// math.Min
	fmt.Println(math.Min(10, 30))
}
