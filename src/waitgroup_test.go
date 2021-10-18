package src

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronously(group *sync.WaitGroup) {
	// tandai selesai (wajib)
	defer group.Done()

	// tambahkan 1 Goroutine
	group.Add(1)

	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronously(group)
	}

	// Wait agar menunggu kode di atas selesai baru execute kode di bawahnya
	group.Wait()
	fmt.Println("Complete!")
}
