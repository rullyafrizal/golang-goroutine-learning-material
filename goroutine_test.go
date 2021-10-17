package main

import (
	"fmt"
	"testing"
	"time"
)

/**
* Goroutine tidak cocok digunakan ketika eksekusi function yang memiliki return value
* bisa tetapi tidak cocok
*
* - Goroutine sangat ringan, jadi kita bisa bikin jutaan goroutine tanpa takut boros memory
*/

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	// menggunakan Goroutine dengan menambahkan keyword "go" dan dilanjut function yang akan diproses menggunakan goroutine
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(2 * time.Second)
}

// Membuat banyak goroutine
func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 1; i <= 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
