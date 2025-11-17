package main

import "fmt"

// Type Declarations
// Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
// Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti

func main() {

	type NoKTP string // alias untuk tipe data string

	var ktpEko NoKTP = "1111111111"

	var contoh string = "222222222"
	var contohKtp NoKTP = NoKTP(contoh) // konversi dari string ke NoKTP dan pada dasarnya NoKTP adalah string

	fmt.Println(ktpEko)
	fmt.Println(contohKtp)

}
