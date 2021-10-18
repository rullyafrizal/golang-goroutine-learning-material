package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var x = 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println("Counter :", x)
}
