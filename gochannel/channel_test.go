package gochannel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

	// membuat channel baru
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		// Memasukkan goroutine(data) ke sebuah channel pakai '<-'
		channel <- "Bayu Firmansyah"
	}()

	// menampung "data" dari channel pakai '<-'
	data := <-channel

	fmt.Println(data)
	close(channel)
}
