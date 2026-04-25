package main

import (
	"encoding/base64"
	"fmt"
)

// Package encoding
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_605#slide=id.g289ef3305ba_0_605
// - Golang menyediakan package encoding untuk melakukan encode dan decode
// - https://pkg.go.dev/encoding
// - Golang menyediakan berbagai macam algoritma untuk encoding, contoh yang populer adalah base64, csv dan json

func main() {
	value := "Eko Kurniawan Khannedy"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}
