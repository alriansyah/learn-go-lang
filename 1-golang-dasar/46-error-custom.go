package main

import "fmt"

// Membuat Custom Error
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g289dd94e9bf_0_77#slide=id.g289dd94e9bf_0_77
// - Karena error adalah sebuah interface, jadi jika kita ingin membuat error sendiri, kita bisa membuat struct yang mengikuti kontrak dari interface error
// - Semua struct dibawah sudah dibuat mengikuti kontrak interface error, karena sudah memiliki method Error() string

// Custom error karena sudah mengikuti kontrak error
type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

// Custom error karena sudah mengikuti kontrak error
type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "eko" {
		return &notFoundError{"data not found"}
	}

	// ok

	return nil
}

func main() {
	err := SaveData("budi", nil)
	if err != nil {
		// terjadi error
		// Cek apakah error bertipe *validationError
		// if validationErr, ok := err.(*validationError); ok {
		// 	// Jika iya, tampilkan pesan error khusus validation
		// 	fmt.Println("validation error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	// Jika bukan, cek apakah error bertipe *notFoundError
		// 	fmt.Println("not found error:", notFoundErr.Error())
		// } else {
		// 	// Jika bukan keduanya, anggap error tidak diketahui
		// 	fmt.Println("unknown error:", err.Error())
		// }

		switch finalError := err.(type) { // melakukan type assertion untuk mendapatkan tipe data dari error
		case *validationError: // case untul tipe pointer validationError
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError: // case untuk tipe pointer notFoundError
			fmt.Println("not found error:", finalError.Error())
		default: // default untuk error yang tidak diketahui
			fmt.Println("unknown error:", finalError.Error())
		}

	} else {
		// sukses
		fmt.Println("Sukses")
	}
}
