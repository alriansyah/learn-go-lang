package helper

import "fmt"

var version = "1.0.0" // private, tidak bisa diakses dari luar package
var Application = "golang" // public, bisa diakses dari luar package

func sayGoodBye(name string) string { // private, tidak bisa diakses dari luar package
	return "Good Bye " + name
}

func SayHello(name string) string { // Pastikan diawali dengan huruf besar agar bisa diakses dari luar package
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Eko")
	fmt.Println(version)
}
