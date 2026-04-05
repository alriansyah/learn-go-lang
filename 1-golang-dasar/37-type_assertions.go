package main

import "fmt"

// Type Assertions
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g83df2f1da9_0_0#slide=id.g83df2f1da9_0_0
// - Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
// - Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong

// Type Assertions Menggunakan Switch
// - Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
// - Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
// - Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions

func random() any { // mengembalikan tipe data any/interface{}
	//return "Eko"
	//return 1
	return true
}

func main() {
	var result any = random()
	//var resultString string = result.(string) // konversi dari any ke string, ini panic/error karerna random() mengembalikan tipe data boolean
	//fmt.Println(resultString)
	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	switch value := result.(type) { // mengambil tipe data dari result
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
