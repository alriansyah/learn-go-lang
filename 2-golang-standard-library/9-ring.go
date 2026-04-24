package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// Package container/ring
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_204#slide=id.g289ef3305ba_0_204
// - Package container/ring adalah implementasi struktur data circular list
// - Circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)
// - https://golang.org/pkg/container/ring/

func main() {
	var data *ring.Ring = ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	//data.Value = "Value 1"
	//
	//data = data.Next()
	//data.Value = "Value 2"
	//
	//data = data.Next()
	//data.Value = "Value 3"
	//
	//data = data.Next()
	//data.Value = "Value 4"
	//
	//data = data.Next()
	//data.Value = "Value 5"

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
