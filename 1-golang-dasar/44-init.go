package main

import (
	"1-golang-dasar/database"
	_ "1-golang-dasar/internal"
	"fmt"
)

// Package Initialization
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g76f41bdaac_0_1100#slide=id.g76f41bdaac_0_1100
// - Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses
// - Ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database
// - Untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init

// Blank Identifier
// - Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package
// - Secara default, Go-Lang akan komplen ketika ada package yang di import namun tidak digunakan
// - Untuk menangani hal tersebut, kita bisa menggunakan blank identifier (_) sebelum nama package ketika melakukan import

func main() {
	fmt.Println(database.GetDatabase())
}
