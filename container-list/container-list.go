package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Moch.")
	data.PushBack("Syahryil")
	data.PushBack("Mahendra")

	data.PushFront("Lucky")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
	fmt.Println(data.Front().Next().Value)
	fmt.Println(data.Back().Prev().Value)
	i := 1
	// menampilkan seluruh data double linked list dari depan
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Print("Data ", i, " = ")
		fmt.Println(e.Value)
		i++
	}

	// menampilkan seluruh data double linked list dari belakang
	i = 4
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Print("Data ", i, " = ")
		fmt.Println(e.Value)
		i--
	}
}
