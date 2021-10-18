package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Cara mengamankan race condition menggunakan Mutex
// Mutual Exclusion
func TestMutexRaceCondition(t *testing.T) {
	var x = 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			for j := 0; j < 100; j++ {
				x += 1
			}
			mutex.Unlock()
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter :", x)
}
