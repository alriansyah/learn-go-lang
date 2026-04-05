package main

import "fmt"

// Interface
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_1166#slide=id.g74a60ad919_0_1166
// - Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
// - Sebuah interface berisikan definisi-definisi method
// - Biasanya interface digunakan sebagai kontrak
// - Interface biasanya diimplementasikan dalam bentuk struct
// - Interface bisa ddisebut dengan prototype atau kontrak

// Interface Kosong
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g76f41bdaac_0_752#slide=id.g76f41bdaac_0_752

// Implementasi Interface
// - Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
// - Sehingga kita tidak perlu mengimplementasikan interface secara manual
// - Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana

// Definisi kontrak interface
type HasName interface {
	// - Ketika ada yang ingin mengikuti intertface ini, maka dia harus memiliki semua yang ada di interface ini
	GetName() string // method GetName dengan return type string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// Person secara umum merupakan implementasi dari interface HasName, karena memiliki method GetName
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Eko"}
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}
