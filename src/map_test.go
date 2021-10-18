package src

import (
	"fmt"
	"sync"
	"testing"
)

// Beberapa function yang bisa dipakai di Map
// Store(k, v) -> menyimpan data di Map with key and value as param
// Load(k) -> mengambil data dari Map by key as param
// Delete(k) -> untuk menghapus data di Map by key as param
// Range(function(k, v)) -> untuk melakukan iterasi seluruh data Map

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, group, i)
	}

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

	group.Wait()
}

func AddToMap(data *sync.Map, group *sync.WaitGroup, val int) {
	defer group.Done()

	group.Add(1)
	data.Store(val, val)
}