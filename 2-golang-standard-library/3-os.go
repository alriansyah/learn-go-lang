package main

import (
	"fmt"
	"os"
)

// Package os
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_91#slide=id.g289ef3305ba_0_91
// - Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
// - Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan  disemua sistem operasi)
// - https://golang.org/pkg/os/

// Contoh cara menjalankan aplkasinya yaitu seperti berikut
// go run 3-os.go argumen1 argumen2

func main() {
	args := os.Args // Mendapatkan argumen yang dikirimkan ketika kita menjalankan aplikasi Go-Lang, misal ketika kita menjalankan go run 3-os.go argumen1 argumen2 maka argumen1 dan argumen2 akan masuk ke dalam os.Args dalam bentuk slice string
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
