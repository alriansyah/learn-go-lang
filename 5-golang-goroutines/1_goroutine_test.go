package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Goroutines tidak cocok untuk function return value
// - Ada masalah ketika ketika menggunakan goroutines untuk function return value, karena return value tidak bisa ditangkap oleh goroutines
func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // running secara asynchronous
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

// Goroutine Sangat Ringan
func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
		// hasilnya tidak akan berurutan, karena goroutines berjalan secara asynchronous
		// - Walaupun dia itu sebenarnya concurrent, dan harusnya berjalan berurutan.
		// Tapi karena laptop yang digunakan multi-core, jadi dia tidak hanya berjalan secara concurrent, tapi dia juga berjalan secara parallel, sehingga hasilnya tidak berurutan.
	}

	time.Sleep(5 * time.Second)
}
