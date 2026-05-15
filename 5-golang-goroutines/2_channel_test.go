package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

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
