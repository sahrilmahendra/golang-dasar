package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Name < value[j].Name
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}
func main() {
	users := []User{
		{"Sahril", 24},
		{"Rudi", 28},
		{"Lucky", 25},
		{"Ahmad", 23},
	}
	fmt.Println("sebelum sort : ", users)

	sort.Sort(UserSlice(users))
	fmt.Println("setelah sort :", users)
}
