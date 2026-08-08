package main

import "fmt"

// Panic
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g74a60ad919_0_954#slide=id.g74a60ad919_0_954
// - Panic function adalah function yang bisa kita gunakan untuk menghentikan program
// - Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
// - Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

// Recover
// - Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// - Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan


// analogy di bahasa pemrogramman lain untuk defer, panic, dan recover adalah try catch, throw, dan finally
// namun di go tidak ada try catch

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("Eko Kurniawan Khannedy")
}
