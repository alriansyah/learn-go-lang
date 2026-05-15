package golang_goroutines

import (
	"fmt"
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
