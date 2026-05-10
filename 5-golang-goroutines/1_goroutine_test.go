package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Goroutines tidak cocok untuk function return value

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(10 * time.Second)
}

// func DisplayNumber(number int) {
// 	fmt.Println("Display", number)
// }

// func TestManyGoroutine(t *testing.T) {
// 	for i := 0; i < 100000; i++ {
// 		go DisplayNumber(i)
// 	}

// 	time.Sleep(5 * time.Second)
// }
