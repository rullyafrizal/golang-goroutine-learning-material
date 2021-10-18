package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Ini adalah simulasi deadlock
// Kode test dibawah ini yang berisi fungsi2 apabila dijalankan akan terjadi deadlock
type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	var user1 = UserBalance{Name: "Rully", Balance: 1_250_000}
	var user2 = UserBalance{Name: "John", Balance: 1_250_000}

	var amount = 100_000

	go Transfer(&user1, &user2, amount)
	go Transfer(&user2, &user1, amount)

	time.Sleep(10 * time.Second)

	fmt.Println("User 1 Balance :", user1.Balance)
	fmt.Println("User 2 Balance :", user2.Balance)

}
