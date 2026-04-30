package main

import "fmt"

// Tipe Data Array
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_1195#slide=id.g74a60ad919_0_1195

func main() {
	var names [3]string

	names[0] = "Eko" // Mengubah nialai di posisi index ke 0
	names[1] = "Kurniawan"
	names[2] = "Khannedy"

	fmt.Println(names)
	fmt.Println(names[0]) // Mengambil nilai di posisi index ke 0
	fmt.Println(names[1])
	fmt.Println(names[2])

	// var values1 = [3]int{100, 80, 90} // Cara lain mendeklarasikan array dengan isi nilai langsung

	// Membuat Array Langsung Dengan Isi Nilai,dengan value didalamnya opsional untuk panjang array
	// menggunakan [...] harus diisi nilainya, jika tidak maka akan error
	var values = [...]int{
		90,
		80,
		95,
		100,
		110,
	}

	// values[5] = 120 // Error karena index ke 5 tidak ada, panjang array values adalah 5 dengan index terakhir 4

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values)) // len() digunakan untuk mengetahui panjang array

	// di go tidaka ada fitur delete pada array, solusinya paling kosongkan nilainya, baik string menjadi "" atau int menjadi 0
}
