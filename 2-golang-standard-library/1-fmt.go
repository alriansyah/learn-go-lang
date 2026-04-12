package main

import "fmt"

// Package fmt
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_577#slide=id.g289ef3305ba_0_577
// - Sebelumnya kita sudah sering menggunakan package fmt dengan menggunakan function Println
// - Selain Println, masih banyak function yang terdapat di package fmt, contohnya banyak digunakan untuk melakukan format
// - https://pkg.go.dev/fmt

func main() {
	firstName := "Eko"
	lastName := "Khannedy"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
