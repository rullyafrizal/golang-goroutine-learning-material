package src

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	// jika transfer ke channel
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <- timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	// langsung akses channel menggunakan After
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <- channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	// jika ingin eksekusi function dengan delay waktu dengan AfterFunc()
	time.AfterFunc(5 * time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}
