package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestHelloName(t *testing.T) {
	go HelloName("Bayu")
	fmt.Println("Ups")

	// tunggu 1 detik
	time.Sleep(1 * time.Second)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i <= 10000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
