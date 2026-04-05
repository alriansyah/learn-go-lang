package main

import "fmt"

// Interface Kosong
// https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit?slide=id.g76f41bdaac_0_752#slide=id.g76f41bdaac_0_752
// - Go-Lang bukanlah bahasa pemrograman yang berorientasi objek
// - Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
// - Contoh di Java ada java.lang.Object
// - Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
// - Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasi nya
// - Interface kosong, juga memiliki type alias bernama any

func Ups() any { // func Ups() interface{} {
	//return 1
	//return true
	return "Ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
