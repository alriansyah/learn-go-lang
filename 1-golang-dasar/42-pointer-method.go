package main

import "fmt"

// Pointer di Method
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g76f41bdaac_0_820#slide=id.g76f41bdaac_0_820
// - Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
// - Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
// - Untuk method sangat direkomendasikan untuk selalu direkomendasikan menggunakan pointer

type Man struct {
	Name string
}

func (man *Man) Married() { // Pointer di Method otomatis dihandle oleh Go
	man.Name = "Mr. " + man.Name
}

func main() {
	eko := Man{"Eko"}
	eko.Married()

	fmt.Println(eko.Name)
}
