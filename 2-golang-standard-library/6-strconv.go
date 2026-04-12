package main

import (
	"fmt"
	"strconv"
)

// Package strconv
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_143#slide=id.g289ef3305ba_0_143
// - Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34
// - Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string, atau sebaliknya
// - Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
// - https://golang.org/pkg/strconv/

// Beberapa Function di Package strconv :
// strconv.parseBool(string) : Mengubah string ke bool
// strconv.parseFloat(string) : Mengubah string ke float
// strconv.parseInt(string) : Mengubah string ke int64
// strconv.FormatBool(bool) : Mengubah bool ke string
// strconv.FormatFloat(float, … ) : Mengubah float64 ke string
// strconv.FormatInt(int, … ) : Mengubah int64 ke string
// strconv.Itoa(int) : Mengubah int ke string
// strconv.Atoi(string) : Mengubah string ke int

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(999)
	fmt.Println(stringInt)
}
