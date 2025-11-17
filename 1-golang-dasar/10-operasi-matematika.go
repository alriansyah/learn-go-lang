package main

import "fmt"

// Operasi Matematika
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_353#slide=id.g74a60ad919_0_353

// Urutan prioritas operasi matematika :
// 1. Tanda Kurung ()
// 2. Perkalian & Pembagian * /
// 3. Penjumlahan & Pengurangan + -

/*
Urutan operasi matematika yang benar (menggunakan aturan BODMAS/PEMDAS):
- Tanda kurung () - operasi dalam kurung dikerjakan terlebih dahulu
- Eksponen/Pangkat ², ³, √ - seperti kuadrat, akar, dll
- Perkalian × dan Pembagian ÷ - memiliki prioritas yang sama, dikerjakan dari kiri ke kanan
- Penjumlahan + dan Pengurangan - - memiliki prioritas yang sama, dikerjakan dari kiri ke kanan
*/

func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = (a + b) - d*e
	fmt.Println(c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)

	var j = 1
	j++ // j = j + 1
	fmt.Println(j)
	j++ // j = j + 1
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)
}
