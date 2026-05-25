package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeout := time.After(5 * time.Second)

	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t)
		case <-timeout:
			return
		}
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	timeout := time.After(5 * time.Second)

	for {
		select {
		case t := <-channel:
			fmt.Println(t)
		case <-timeout:
			return
		}
	}
}
