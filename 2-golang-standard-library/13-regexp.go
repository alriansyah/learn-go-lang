package main

import (
	"fmt"
	"regexp"
)

// Package regexp
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_291#slide=id.g289ef3305ba_0_291
// - Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
// - Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
// - https://github.com/google/re2/wiki/Syntax
// - https://golang.org/pkg/regexp/

// Beberapa Function di Package regexp :
// regexp.MustCompile(string) : Membuat Regexp
// Regexp.MatchString(string) bool : Mengecek apakah Regexp match dengan string
// Regexp.FindAllString(string, max) : Mencari string yang match dengan maximum jumlah hasil, retrnnya berupa slice of string

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eKo"))

	fmt.Println(regex.FindAllString("eko edo egi ego e1o eto eKo", 10))
}
