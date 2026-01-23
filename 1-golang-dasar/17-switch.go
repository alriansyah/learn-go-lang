package main

import "fmt"

// Switch Expression
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_474#slide=id.g74a60ad919_0_474

func main() {
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// Switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Switch tanpa kondisi/ekspresi
	// - Kondisi di switch expression tidak wajib
	// - Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap case nya
	// - kondisi seperti ini disarankan menggunakan if-else saja
	name = "KhannedyKhannedyKhannedy"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
