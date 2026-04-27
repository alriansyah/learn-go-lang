package main

import (
	"fmt"
	"path"
)

// Package path
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_617#slide=id.g289ef3305ba_0_617
// - Package path digunakan untuk memanipulasi data path seperti path di URL atau path di File System
// - Secara default Package path menggunakan slash sebagai karakter path nya, oleh karena itu cocok untuk data URL
// - https://pkg.go.dev/path
// - Namun jika ingin menggunakan untuk memanipulasi path di File System, karena Windows menggunakan backslash, maka khusus untuk File System, perlu menggunakan pacakge path/filepath
// - https://pkg.go.dev/path/filepath

func main() {
	fmt.Println(path.Dir("hello/world.go"))             // Mengambil directory dari path, return string
	fmt.Println(path.Base("hello/world.go"))            // Mengambil nama file dari path, return string
	fmt.Println(path.Ext("hello/world.go"))             // Mengambil ekstensi file dari path, return string
	fmt.Println(path.Join("hello", "world", "main.go")) // Menggabungkan beberapa bagian path, return string
}
