package src

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = &sync.Mutex{}
var cond = sync.NewCond(locker)
var group = sync.WaitGroup{}

func TestCond(t *testing.T) {
	for i := 1; i <= 10; i++ {
		go WaitCondition(i)
	}

	// jika menggunakan signal -> beritahu sebuah Goroutine saaja
	//go func() {
	//	for i := 1; i <= 10; i++ {
	//		time.Sleep(1 * time.Second)
	//		cond.Signal()
	//	}
	//}()

	// jika menggunakan broadcast -> beritahu semua Goroutine secara langsung
	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}

func WaitCondition(val int) {
	defer group.Done()
	group.Add(1)

	// lock condition terlebih dahulu
	cond.L.Lock()

	// tunggu dulu boleh jalan atau tidak setelah berhasil locking
	cond.Wait()

	fmt.Println("Done", val)

	// Unlock
	cond.L.Unlock()

}