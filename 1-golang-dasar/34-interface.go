package main

import "fmt"

// Interface
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_1166#slide=id.g74a60ad919_0_1166

// Definisi kontrak interface
type HasName interface {
	GetName() string
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
