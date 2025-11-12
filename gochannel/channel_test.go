package gochannel

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func giveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	//memasukkan data ke parameter channel
	channel <- "Bayu Firmansyah"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go giveMeResponse(channel)
	//menangkap data di channel
	data := <-channel

	fmt.Println(data)
}

// chan <- {type data} untuk memasukkan data ke channel saja (in only)
func inOnly(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Bayu Firmansyah"
}

// {data} <- chan {type data} untuk memasukkan data dari channel (out only)
func outOnly(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInAndOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go inOnly(channel)
	go outOnly(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	//3 merupakan capasitas yang dimiliki oleh channel
	//berfungsi agar channel dapat menampung sampai 3 data sebelum di ambil oleh goroutine lain
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Bayu"
		channel <- "Firmansyah"
		channel <- "Silfi"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Data ke-" + strconv.Itoa(i)
		}

		close(channel)
	}()

	//range dapat digunakan ketika kita tidak tahu seberapa banyak data yang diterima oleh sebuah channel
	for data := range channel {
		fmt.Println(data)
	}

}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go giveMeResponse(channel1)
	go giveMeResponse(channel2)

	counter := 0
	for {
		//select(seperti switch case) akan memilih channel tercepat yang mempunyai data untuk diambil
		select {
		case data := <-channel1:
			fmt.Println("Data pada channel ke-1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data pada channel ke-2", data)
			counter++
		//	default berfungsi ketika channel masih kosong, maka default yang akan dijalankan
		default:
			fmt.Println("Menunggu Data...")
		}

		if counter == 2 {
			break
		}
	}
}

func TestChannel(t *testing.T) {

	// membuat channel baru
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		// Memasukkan goroutine(data) ke sebuah channel pakai '<-'
		channel <- "Bayu Firmansyah"
	}()

	// menampung "data" dari channel pakai '<-'
	data := <-channel

	fmt.Println(data)
}
