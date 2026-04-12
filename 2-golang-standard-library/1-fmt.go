package main

import "fmt"

// Package fmt
// - Sebelumnya kita sudah sering menggunakan package fmt dengan menggunakan function Println
// - Selain Println, masih banyak function yang terdapat di package fmt, contohnya banyak digunakan untuk melakukan format
// - https://pkg.go.dev/fmt

func main() {
	firstName := "Eko"
	lastName := "Khannedy"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
