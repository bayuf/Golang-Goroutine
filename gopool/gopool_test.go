package gopool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// pool merupakan wadah data yang reusable
func TestGoPool(t *testing.T) {
	pool := sync.Pool{ // membuat pool
		New: func() any {
			return "new"
		},
	}

	pool.Put("Bayu") // menambah data ke dalam pool
	pool.Put("Firmansyah")
	pool.Put("Silfi Dian Putri")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get() // ambil data
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data) // mengembalikan data setelah di ambil
		}()

	}

	time.Sleep(10 * time.Second)
}
