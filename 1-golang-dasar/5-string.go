package main

import "fmt"

// Tipe Data String
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_198#slide=id.g74a60ad919_0_198

// Function untuk string
// len(“string”): Menghitung jumlah karakter di String (spasi/tab/enter dihitung)
// “string”[number]: Mengambil karakter pada posisi yang ditentukan

func main() {
	fmt.Println("Eko")
	fmt.Println("Eko Kurniawan")
	fmt.Println("Eko Kurniawan Khannedy")

	fmt.Println(len("Eko"))
	fmt.Println("Eko Kurniawan"[0]) // mengambil karakter pada index ke 0: hasilnya 69 (type byte)
	fmt.Println("Eko Kurniawan Khannedy"[1]) // mengambil karakter pada index ke 1: hasilnya 107 (type byte)
}
