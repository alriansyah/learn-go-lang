package golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Membuat Channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Eko Kurniawan Khannedy" // mengirim data dari channel
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel // menerima data dari channel, jika channel belum ada data, maka akan menunggu sampai ada data yang dikirim ke channel
	// fmt.Println(<-channel) // seperti ini juga bisa
	fmt.Println(data)

	time.Sleep(5 * time.Second)

	// kenapa data := <-channel tidak menerima goroutine :
	// - karena saat menjalankan sebuah program sebenarnya dia running as a goroutine, cuma kita kita ngga sadar aja kalau itu running sebagai goroutine
}

// Channel sebagain parameter
func GiveMeResponse(channel chan string) { // ketika menggunakan channel sebagai parameter, kita tidak perlu menggunakan pointer
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan Khannedy"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// Channel In dan Out
func OnlyIn(channel chan<- string) { // Mengirim data ke channel
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan Khannedy"
}

func OnlyOut(channel <-chan string) { // Menerima data dari channel
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered Channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Eko"
		channel <- "Kurniawan"
		channel <- "Khannedy"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

// Range Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}
