package main

import (
	"container/list"
	"fmt"
)

// Package container/list
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_181#slide=id.g289ef3305ba_0_181
// - Package container/list adalah implementasi struktur data double linked list di Go-Lang
// - https://golang.org/pkg/container/list/
// - Fitur ini biasanya digunakan untuk antrian

func main() {
	var data *list.List = list.New()

	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // eko

	next := head.Next() // kurniawan
	fmt.Println(next.Value)

	next = next.Next() // khannedy
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
