package main

import "fmt"

// Function Return Value
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_649#slide=id.g74a60ad919_0_649

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Eko")
	fmt.Println(result)

	fmt.Println(getHello("Budi"))
	fmt.Println(getHello("Joko"))
}
