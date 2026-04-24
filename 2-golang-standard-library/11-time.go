package main

import (
	"fmt"
	"time"
)

// Package time
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_248#slide=id.g289ef3305ba_0_248
// - Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
// - https://golang.org/pkg/time/

// Beberapa Function di Package time :
// time.Now() : Untuk mendapatkan waktu saat ini
// time.Date(...) : Untuk membuat waktu
// time.Parse(layout, string) : Untuk memparsing waktu dari string

// Duration
// - Saat menggunakan tipe data waktu, kadang kita butuh data berupa durasi
// - Package tipe memiliki type Duration, yang sebenarnya adalah alias untuk int64
// - Namun terdapat banyak method yang bisa kita gunakan untuk memanipulasi Duration

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)

	var utc time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2020-10-10 10:10:10"
	// value := "ASAL"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
}
