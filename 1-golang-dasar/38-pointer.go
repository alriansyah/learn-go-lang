package main

import "fmt"

// Pointer
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g83df2f1da9_0_12#slide=id.g83df2f1da9_0_12

type Address struct {
	City, Province, Country string
}

func main() {
	//address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	//address2 := address1 // copy value
	//
	//address2.City = "Bandung"
	//fmt.Println(address1) // tidak berubah
	//fmt.Println(address2) // berubah menjadi Bandung

	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer, pointer ke address1

	address2.City = "Bandung"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Bandung
}
