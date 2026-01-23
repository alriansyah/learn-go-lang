package main

import "fmt"

// For Loops
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_513#slide=id.g74a60ad919_0_513

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// for dengan statement init dan post (init; condition; post)
	// step awal melakukan init (dieksesi pertama kali dan hanya sekali)
	// kemudian dicek kondisinya, jika true maka blok kode di dalam for akan dijalankan, setelah itu melakukan post, lalu kembali ke cek kondisi dan seterusnya hingga kondisi bernilai false
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	fmt.Println("Selesai")

	// For Range
	// - For bisa digunakan untuk melakukan iterasi terhadap semua data collection
	// - Data collection contohnya Array, Slice dan Map

	// ini cara manual
	names := []string{"Eko", "Kurniawan", "Khannedy"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// ini cara menggunakan for range (index/key, value)
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	// jika tidak ingin menggunakan index/key, bisa menggunakan underscore _
	for _, name := range names {
		fmt.Println(name)
	}
}
