package src

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}


	totalCpu := runtime.NumCPU()
	fmt.Println("CPU :", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine :", totalGoroutine) // default 2 : 1 untuk execute program, dan 1 untuk garbage collector

	group.Wait()
}
