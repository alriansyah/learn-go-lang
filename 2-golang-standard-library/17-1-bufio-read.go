package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Package io
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_621#slide=id.g289ef3305ba_0_621
// - IO atau singkatan dari Input Output, merupakan fitur di Golang yang digunakan sebagai standard untuk proses Input Output
// - Di Golang, semua mekanisme input output pasti mengikuti standard package io
// - https://pkg.go.dev/io

// Implementasi IO
// - Penggunaan dari IO sendiri di Golang terdapat dibanyak package, sebelumnya contohnya kita menggunakan CSV Reader dan CSV Writer
// - Karena Package IO sebenarnya hanya kontrak untuk IO, untuk implementasinya kita harus lakukan sendiri
// - Tapi untungnya, Golang juga menyediakan package untuk mengimplementasikan IO secara mudah, yaitu menggunakan package bufio

// Package bufio
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_310#slide=id.g289ef3305ba_0_310
// - Package bufio atau singkatan dari buffered io
// - Package ini digunakan untuk membuat data IO seperti Reader dan Writer
// - https://pkg.go.dev/bufio

func main() {
	input := strings.NewReader("this is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
