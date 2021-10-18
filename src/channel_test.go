package src

import (
	"fmt"
	"strconv"
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


// TestBufferedChannel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	// tidak akan error meski tidak ada Goroutine yang mengambil data
	// karena data ditampung di buffer
	go func() {
		channel <- "Rully Afrizal"
		channel <- "Ardelia Pramesti"
	}()

	go func() {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("selesai")
	fmt.Println(cap(channel)) // Panjang buffer
	fmt.Println(len(channel)) // Banyak data yang ada di buffer

}


// TestRangedChannel
func TestRangedChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- "Perulangan ke - " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data :", data)
	}

	fmt.Println("Selesai")
}


func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1 :", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2 :", data)
			counter++
		// Default select = apabila di kedua channel tidak ada datanya
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}

	}
}