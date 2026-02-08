package main

import "fmt"

// Type Assertions
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g83df2f1da9_0_0#slide=id.g83df2f1da9_0_0

func random() any {
	return true
}

func main() {
	var result any = random()
	//var resultString string = result.(string) // konversi dari any ke string, ini panic/error karerna random() mengembalikan tipe data boolean
	//fmt.Println(resultString)
	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	switch value := result.(type) {  // mengambil tipe data dari result
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
