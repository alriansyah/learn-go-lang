package main

import "fmt"

// Continue
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_560#slide=id.g74a60ad919_0_560
// - Break & continue adalah kata kunci yang bisa digunakan dalam perulangan
// - Break digunakan untuk menghentikan seluruh perulangan
// - Continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
