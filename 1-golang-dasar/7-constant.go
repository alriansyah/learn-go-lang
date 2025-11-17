package main

import "fmt"

// Constant
// Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
// Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
// Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya
// Ketika variable constant telah dibuat, variable tersebut jika tidak dipakai maka tidak akan error seperti di variable biasa

// Deklarasi Multiple Constant
// Sama seperti variable, di Go-Lang  juga kita bisa membuat constant secara sekaligus banyak

func main() {
	// const firstName string = "Eko"
	// const lastName = "Khannedy"
	const (
		firstName string = "Eko"
		lastName         = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Budi"
	// lastName = "Joko"
}
