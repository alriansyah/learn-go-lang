package main

import "fmt"

// Recursive Function (Function memanggil dirinya sendiri)
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g76f41bdaac_0_641#slide=id.g76f41bdaac_0_641

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(result)
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}
