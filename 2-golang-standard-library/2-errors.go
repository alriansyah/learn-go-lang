package main

import (
	"errors"
	"fmt"
)

// Package errors
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_640#slide=id.g289ef3305ba_0_640
// - Sebelumnya kita sudah membahas tentang interface error yang merupakan representasi dari error di Go-Lang, dan membuat error menggunakan function errors.New()
// - Sebenarnya masih banyak yang bisa kita lakukan menggunakan package errors, contohnya ketika kita ingin membuat beberapa value error yang berbeda
// - https://pkg.go.dev/errors

// Mengecek Jenis Error
// - Misal kita membuat jenis error sendiri, lalu kita ingin mengecek jenis errornya
// - Kita bisa menggunakan errors.Is() untuk mengecek jenis type error nya

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "eko" {
		return NotFoundError
	}

	// sukses

	return nil
}

func main() {
	err := GetById("eko")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
