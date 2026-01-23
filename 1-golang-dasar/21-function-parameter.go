package main

import "fmt"

// Function Parameter
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_593#slide=id.g74a60ad919_0_593

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Eko", "Khannedy")
	sayHelloTo("Budi", "Nugraha")
}
