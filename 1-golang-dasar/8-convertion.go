package main

import "fmt"

// Konversi Tipe Data
// Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
// Misal kita ingin mengkonversi tipe data int32 ke int63, dan lain-lain

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) // kelebihan kapasitas/number overflow, jadi nilainya akan berubah menjadi negatif (kembali ke belakang)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Eko Kurniawan"
	var e uint8 = name[0] // otomatis menjadi dalam bentuk byte/uint8, sehingga perlu di konversi ke string
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
