package main

import "fmt"

// Break
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_560#slide=id.g74a60ad919_0_560

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}
}
