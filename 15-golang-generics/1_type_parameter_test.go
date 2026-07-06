package golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	// var result string = Length[string]("Eko") // go version 1.18
	var result string = Length("Eko")
	assert.Equal(t, "Eko", result)

	// var resultNumber int = Length[int](100) // go version 1.18
	var resultNumber int = Length(100)
	assert.Equal(t, 100, resultNumber)
}
