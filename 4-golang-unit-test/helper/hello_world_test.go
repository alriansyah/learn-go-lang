package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		// error
		t.Error("Result must be 'Hello Eko'")
	}

	fmt.Println("TestHelloWorldEko Done")
}

func TestHelloWorldRian(t *testing.T) {
	result := HelloWorld("Rian")

	if result != "Hello Rian" {
		// error
		t.Fatal("Result must be 'Hello Rian'")
	}

	fmt.Println("TestHelloWorldRian Done")
}
