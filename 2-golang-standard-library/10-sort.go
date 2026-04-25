package main

import (
	"fmt"
	"sort"
)

// Package sort
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_219#slide=id.g289ef3305ba_0_219
// - Package sort adalah package yang berisikan utilitas untuk proses pengurutan
// - Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
// - https://golang.org/pkg/sort/

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int { // Menentukan jumlah panjang data
	return len(s)
}

func (s UserSlice) Less(i, j int) bool { // Menentukan aturan pengurutan data, jika i < j maka akan diurutkan dari kecil ke besar, jika i > j maka akan diurutkan dari besar ke kecil
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	// Menukar posisi data pada index i dan j dalam slice
	// Digunakan oleh package sort saat proses pengurutan
	s[i], s[j] = s[j], s[i]
} 

func main() {
	users := []User{
		{"Eko", 30},
		{"Budi", 35},
		{"Joko", 25},
		{"Adit", 20},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
