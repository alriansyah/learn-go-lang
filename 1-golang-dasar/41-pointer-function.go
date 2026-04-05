package main

import "fmt"

// Pointer di Function
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g76f41bdaac_0_1011#slide=id.g76f41bdaac_0_1011
// - Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
// - Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
// - Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
// - Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
// - Untuk melakukan ini, kita juga bisa menggunakan pointer di function
// - Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address Address = Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
