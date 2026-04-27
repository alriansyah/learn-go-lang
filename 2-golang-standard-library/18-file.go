package main

import (
	"bufio"
	"io"
	"os"
	// "fmt"
)

// File Manipulation
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_835#slide=id.g289ef3305ba_0_835

// File Management
// - Di package os, terdapat File Management, namun sengaja ditunda pembahasannya, karena kita harus tahu dulu tentang IO
// - Saat kita membuat atau membaca file menggunakan Package os, struct File merupakan implementasi dari io.Reader dan io.Writer
// - Oleh karena itu, kita bisa melakukan baca dan tulis terhadap File tersebut menggunakan Package io / bufio

// Open File
// - Untuk membuat / membaca File, kita bisa menggunakan os.OpenFile(name, flag, permission)
// - name berisikan nama file, bisa absolute atau relative / local
// - flag merupakan penanda file, apakah untuk membaca, menulis, dan lain-lain
// - permission, merupakan permission yang diperlukan ketika membuat file, bisa kita simulasikan disini : https://chmod-calculator.com/

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close() // untuk memastikan file akan tertutup setelah selesai digunakan
	file.WriteString(message)
	return nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	// createNewFile("sample.log", "this is sample log") // untuk membuat file baru dengan nama sample.log dan menulis "this is sample log" ke dalamnya

	// result, _ := readFile("sample.log") // untuk membaca file
	// fmt.Println(result)

	addToFile("sample.log", "\nthis is add message") // membuka file sample.log, lalu menambahkan "this is add message" ke dalamnya
}
