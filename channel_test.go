package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// membuat channel dengan make(chan {tipe-data})
	var channel = make(chan string)
	// menutup channel
	defer close(channel)


	//// mengirim data ke channel
	//channel <- "Rully Afrizal"
	//
	//// menerima data dari channel
	//var nama string
	//nama = <- channel
	//
	//fmt.Println(nama)

	go func() {
		time.Sleep(2 * time.Second)
		// Goroutine akan blocked apabila belum ada yang mengambil dari channel
		channel <- "Rully Afrizal Alwin"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}


// Channel sebagai parameter
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	defer close(channel)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Data : Rully Afrizal Alwin"
}


// Channel in and out
// Menandai apakah channel di parameter digunakan untuk mengirim atau menerima data saja

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
	close(channel)
}

// Hanya memasukkan data ke dalam channel
func OnlyIn (channel chan<- string) {
	time.Sleep(2 * time.Second)
	//data := <- channel // error jika dipaksa untuk mengeluarkan data dari dalam channel
	channel <- "Data dikirim ke channel"
}

// Hanya mengeluarkan data dari dalam channel
func OnlyOut(channel <-chan string) {
	data := <- channel
	fmt.Println(data)
}
