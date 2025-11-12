package go_race_condition

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	//mutex digunakan agar tidak terjadi sebuah race condition
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				//lock untuk mengunci proses agar tidak bisa di akses oleh goroutine lain
				mutex.Lock()
				x = x + 1
				//unlock untuk membuka kunci setelah goroutine menyelesaikan tugasnya
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	balance int
}

func (account *BankAccount) addBalance(amount int) {
	//lock untuk RWMutex Write
	account.RWMutex.Lock()
	account.balance += amount
	account.RWMutex.Unlock()

}

func (account *BankAccount) getBalance() int {
	//Rlock untuk RWMutex Read
	account.RWMutex.RLock()
	balance := account.balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.addBalance(1)
				fmt.Println(account.getBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total :", account.getBalance())
}
