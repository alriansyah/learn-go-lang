package main

import "fmt"

// Variable
// Varibale yang sudah dideklarasikan bisa diubah nilainya namun tidak bisa dibuat ulang dengan nama yang sama dalam satu scope
// Variable yang telah di deklarasikan namun tidak digunakan akan menyebabkan error
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_237#slide=id.g74a60ad919_0_237

func main() {
	// var name string
	// var name = "Eko Kurniawan" // dengan inisialisasi otomatisstring, jadi tidak wajinb menyebutkan tipe data
	name := "Eko Kurniawan" // cara singkat deklarasi variable tanpa var
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	// Deklarasi multiple variable
	var (
		firstName  = "Eko"
		middleName = "Kurniawan"
		lastName   = "Khannedy"
		noName string
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
	fmt.Println(noName)
}
