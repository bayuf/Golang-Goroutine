package goonlyonce

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func TestOnlyOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		group.Add(1)
		go once.Do(onlyonce) // menjalankan function sekali di goroutine
		group.Done()
	}
	group.Wait()
	fmt.Println(counter)
}

func onlyonce() {
	counter++
}
