package main

import "fmt"

// Struct
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_1006#slide=id.g74a60ad919_0_1006

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eko Customer // object eko dari struct Customer
	fmt.Println(eko)

	eko.Name = "Eko Kurniawan"
	eko.Address = "Indonesia"
	eko.Age = 30
	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	// Struct Literal
	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)
}
