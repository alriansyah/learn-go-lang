package main

import (
	"fmt"
	"strings"
)

// Package strings
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_124#slide=id.g289ef3305ba_0_124
// - Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String
// - Ada banyak sekali function yang bisa kita gunakan
// - https://golang.org/pkg/strings/

// Beberapa Function di Package strings :
// strings.Trim(string, cutset) : Memotong cutset di awal dan akhir string
// strings.ToLower(string) : Membuat semua karakter string menjadi lower case
// strings.ToUpper(string) : Membuat semua karakter string menjadi upper case
// strings.Split(string, separator) : Memotong string berdasarkan separator
// strings.Contains(string, search) : Mengecek apakah string mengandung string lain
// strings.ReplaceAll(string, from, to) : Mengubah semua string dari from ke to

func main() {
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Split("Eko Kurniawan", " ")) // Outputnya akan menjadi slice string
	fmt.Println(strings.ToLower("Eko Kurniawan"))
	fmt.Println(strings.ToUpper("Eko Kurniawan"))
	fmt.Println(strings.Trim("      Eko Kurniawan      ", " "))
	fmt.Println(strings.ReplaceAll("Eko Kurniawan Eko Khannedy", "Eko", "Budi"))
}
