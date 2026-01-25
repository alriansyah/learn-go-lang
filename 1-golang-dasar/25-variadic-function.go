package main

import "fmt"

// Variadic Function
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g76f41bdaac_0_756#slide=id.g76f41bdaac_0_756

func sumAll(numbers ...int) int { // ... bisa disebut varargs (variable arguments), yang otomatis akan diubah menjadi slice
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10, 10} // argumen slice yang dikirimkan ke parameter variadic harus ditambahkan ... agar menjadi varargs
	fmt.Println(sumAll(numbers...))
}
