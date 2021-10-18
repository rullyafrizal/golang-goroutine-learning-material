package src

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool sync.Pool
	pool = sync.Pool{
		New: func() interface{} {
			return "New" // Default attribute dari pool
		},
	}
	group := &sync.WaitGroup{}

	pool.Put("Rully")
	pool.Put("Afrizal")
	pool.Put("Alwin")

	for i := 0; i <10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	group.Wait()
	time.Sleep(11 * time.Second)
	fmt.Println("Done!")
}
