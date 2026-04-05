package helper

import "fmt"

var version = "1.0.0"
var Application = "golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string { // Pastikan diawali dengan huruf besar agar bisa diakses dari luar package
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Eko")
	fmt.Println(version)
}
