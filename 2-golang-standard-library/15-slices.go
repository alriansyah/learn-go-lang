package main

import (
	"fmt"
	"slices"
)

// Package slices
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_625#slide=id.g289ef3305ba_0_625
// - Di Golang versi terbaru, terdapat fitur bernama Generic, fitur ini akan kita bahas khusus dikelas Golang Generic
// - Fitur Generic ini membuat kita bisa membuat parameter dengan tipe yang bisa berubah-ubah, tanpa harus menggunakan interface kosong / any
// - Salah satu package yang menggunakan fitur Generic ini adalah package slices
// - Package slices ini digunakan untuk memanipulasi data di slice
// - https://pkg.go.dev/slices

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names))             // Min : untuk mencari nilai minimum dari slice, return string atau int tergantung tipe data slice
	fmt.Println(slices.Min(values))            // Min : untuk mencari nilai minimum dari slice, return string atau int tergantung tipe data slice
	fmt.Println(slices.Max(names))             // Max : untuk mencari nilai maksimum dari slice, return string atau int tergantung tipe data slice
	fmt.Println(slices.Max(values))            // Max : untuk mencari nilai maksimum dari slice, return string atau int tergantung tipe data slice
	fmt.Println(slices.Contains(names, "Eko")) // Contains : untuk mengecek apakah data ada di dalam slice atau tidak, return bool
	fmt.Println(slices.Index(names, "Eko"))    // Index : untuk mengecek index dari data di dalam slice, return int
	fmt.Println(slices.Index(names, "Paul"))   // Index : untuk mengecek index dari data di dalam slice, return int
}
