package main

import "fmt"

// Struct
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_1006#slide=id.g74a60ad919_0_1006
// - Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
// - Struct biasanya representasi data dalam program aplikasi yang kita buat
// - Data di struct disimpan dalam field
// - Sederhananya struct adalah kumpulan dari field

// Struct Literal
// - Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct

// Struct Method
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_1162#slide=id.g74a60ad919_0_1162
// - Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
// - Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
// - Method adalah function

type Customer struct {
	Name, Address string
	Age           int
}

// Struct Method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
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

	budi.sayHello("Agus")
	eko.sayHello("Agus")
	joko.sayHello("Agus")
}
