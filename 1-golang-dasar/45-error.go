package main

import (
	"errors"
	"fmt"
)

// Error
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g289dd94e9bf_0_52#slide=id.g289dd94e9bf_0_52

// error Interface
// - Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error

// Membuat Error
// - Untuk membuat error, kita tidak perlu manual.
// - Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (Package sudah di bahas secara detail di materi sebelumnya)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
