package test

import (
	"fmt"
	"golang-dependency-injection/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}
