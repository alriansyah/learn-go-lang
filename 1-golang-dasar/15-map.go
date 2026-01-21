package main

import "fmt"

// Tipe Data Map
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_1238#slide=id.g74a60ad919_0_1238

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Eko"
	//person["address"] = "Subang"

	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	// fmt.Println(person["rawrr"]) // akan menghasilkan nilai kosong jika key tidak ditemukan
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Eko"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
