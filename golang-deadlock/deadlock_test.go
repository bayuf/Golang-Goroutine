package golangdeadlock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) UnLock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(fromUser, toUser *UserBalance, amount int) {
	fromUser.Lock()
	fmt.Println("Lock userFrom:", fromUser.Name)
	fromUser.Change(-amount)

	time.Sleep(1 * time.Second)

	toUser.Lock()
	fmt.Println("Lock userTo: ", toUser.Name)
	toUser.Change(amount)

	time.Sleep(1 * time.Second)

	fromUser.UnLock()
	toUser.UnLock()

}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Bayu",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Silfi",
		Balance: 1000000,
	}

	// Perintah ini akan menyebabkan deadlock karena kedua user sama" di lock oleh Mutex
	// dan Mutex sama" menunggu untuk Unlock dari keduanya
	go Transfer(&user1, &user2, 500000)
	go Transfer(&user2, &user1, 250000)

	time.Sleep(3 * time.Second)

	fmt.Println("Uang Bayu:", user1.Balance)
	fmt.Println("Uang Silfi:", user2.Balance)

}
