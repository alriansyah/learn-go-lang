package main

import "fmt"

// Operator new
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g76f41bdaac_0_1000#slide=id.g76f41bdaac_0_1000
// - Sebelumnya untuk membuat pointer dengan menggunakan operator &
// - Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
// - Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Address = new(Address) // &Address{}
	var alamat2 *Address = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
