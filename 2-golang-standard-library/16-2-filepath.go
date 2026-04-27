package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/world.go"))             // Mengambil directory dari path, return string
	fmt.Println(filepath.Base("hello/world.go"))            // Mengambil nama file dari path, return string
	fmt.Println(filepath.Ext("hello/world.go"))             // Mengambil ekstensi file dari path, return string
	fmt.Println(filepath.IsAbs("hello/world.go"))           // Mengecek apakah path adalah absolute, return bool
	fmt.Println(filepath.IsLocal("hello/world.go"))         // Mengecek apakah path adalah local, return bool
	fmt.Println(filepath.Join("hello", "world", "main.go")) // Menggabungkan beberapa bagian path, return string
}
