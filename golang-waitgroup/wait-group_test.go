package golangwaitgroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	// dilakukan agar add(-1)
	defer group.Done()

	// jumlah proses yang ditunggu
	group.Add(1)
	fmt.Println("Hallo Dunia")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	// waitgroup digunakan untuk menunggu semua proses goroutine selesai dilakukan
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	//tunggu sampai goroutine selesai
	group.Wait()
	fmt.Println("Go-Routine Selesai")
}
