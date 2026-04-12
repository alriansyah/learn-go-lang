package main

import "fmt"

func pointerDiFunction(x *int) {
	*x = 100 // ubah nilai yang ditunjuk oleh pointer

	// Penjelasan:
	// *int → parameter pointer
	// &x → kirim alamat
	// *x = 100 → ubah nilai asli
}

type User struct {
	Name string
}

// Pointer di method
func (u *User) ChangeName(name string) {
	u.Name = name // ubah nilai field Name dari struct User
}

func main() {
	x := 10
	var p *int = &x // p := &x juga bisa digunakan untuk mendeklarasikan pointer

	fmt.Println("Nilai x:", x)
	fmt.Println("Alamat x:", &x)
	fmt.Println("Isi pointer p:", *p)

	*p = 10        // ubah nilai x lewat pointer
	fmt.Println(x) // output: 10

	pointerDiFunction(&x)
	fmt.Println("Nilai x:", x)

	user := User{Name: "Budi"}
	user.ChangeName("Andi")

	fmt.Println(user.Name) // "Andi" ✅
}

// Pointer adalah variabel yang menyimpan alamat memori dari variabel lain, bukan nilainya langsung. pass by reference

// Asterisk Operator (*)
// - Di Golang, * punya 2 fungsi berbeda tergantung konteks:
// 1. Pointer (Sebagai tipe pointer)
// var p *int // p adalah pointer ke tipe int

// 2. Sebagai Dereference (ambil nilai)
// fmt.Println(*p) // ambil nilai dari alamat yang disimpan di pointer
