package main

import (
	"1-golang-dasar/helper"
	"fmt"
)

// Package dan Import
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g76f41bdaac_0_533#slide=id.g76f41bdaac_0_533

// Package
// - Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di Go-Lang
// - Dengan menggunakan package, kita bisa merapikan kode program yang kita buat
// -Package sendiri sebenarnya hanya direktori folder di sistem operasi kita

// Import
// - Secara standar, file Go-Lang hanya bisa mengakses file Go-Lang lainnya yang berada dalam package yang sama
// - Jika kita ingin mengakses file Go-Lang yang berada diluar package, maka kita bisa menggunakan Import

// Access Modifier
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g83df2f1da9_0_4#slide=id.g83df2f1da9_0_4
// - Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable
// - Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable
// - Jika nama nya diawali dengan hurup besar, maka artinya bisa diakses dari package lain, jika dimulai dengan hurup kecil, artinya tidak bisa diakses dari package lain

func main() {
	result := helper.SayHello("Eko")
	fmt.Println(result)

	fmt.Println(helper.Application)
	helper.Contoh()
	// fmt.Println(helper.version) // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Eko")) // tidak bisa diakses
}
