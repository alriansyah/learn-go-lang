package main

import (
	"flag"
	"fmt"
)

// Package flag
// https://docs.google.com/presentation/d/1w5QbuY3HkWYeDrNVkf1s-xuxxsn50mFYGiKI1a6hfMs/edit?slide=id.g289ef3305ba_0_110#slide=id.g289ef3305ba_0_110
// - Package flag berisikan fungsionalitas untuk memparsing command line argument
// - https://golang.org/pkg/flag/

// Contoh cara menjalankan aplkasinya yaitu seperti berikut
// go run 4-flag.go -username=eko -password="rahasia banget" -host=localhost -port=3306

func main() {
	// (nama flag, default value, deskripsi)
	var username *string = flag.String("username", "root", "database username") // return value dari flag.String() adalah pointer
	var password *string = flag.String("password", "root", "database password") // return value dari flag.String() adalah pointer
	var host *string = flag.String("host", "localhost", "database host")        // return value dari flag.String() adalah pointer
	var port *int = flag.Int("port", 0, "database port")                        // return value dari flag.Int() adalah pointer

	flag.Parse() // Melakukan parsing terhadap command line argument, harus dipanggil setelah mendefinisikan flag nya, karena jika tidak dipanggil maka flag nya tidak akan terbaca

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)
}
