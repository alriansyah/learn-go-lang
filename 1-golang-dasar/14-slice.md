# Memahami Slice di Golang (Untuk Pemula)

## 1. Apa itu Slice?

Slice adalah struktur data di Go yang digunakan untuk menyimpan **kumpulan data dengan tipe yang sama**, mirip seperti array, tapi jauh lebih fleksibel karena **ukurannya bisa berubah-ubah (dinamis)**.

Kalau kamu pernah pakai bahasa lain seperti JavaScript (`Array`) atau Python (`List`), slice di Go itu konsepnya mirip — bedanya slice punya tipe data yang tetap (misalnya cuma bisa isi `int` semua, atau `string` semua).

### Kenapa bukan Array saja?

Di Go, **array** punya ukuran tetap yang harus ditentukan di awal dan tidak bisa diubah:

```go
var angka [5]int // hanya bisa menampung 5 elemen, selamanya
```

Sedangkan **slice** dibuat di atas array, tapi ukurannya bisa bertambah/berkurang secara otomatis. Karena itu, di dunia nyata, hampir semua orang Go lebih sering pakai slice daripada array.

---

## 2. Cara Membuat Slice

Ada beberapa cara membuat slice:

### a. Slice Literal (langsung isi nilai)

```go
buah := []string{"Apel", "Jeruk", "Mangga"}
```

### b. Menggunakan `make()`

```go
angka := make([]int, 3) // slice int dengan panjang 3, isi awal: [0 0 0]
```

### c. Slice kosong (nil slice)

```go
var kosong []int // panjang 0, belum ada data
```

---

## 3. Anatomi Slice: Length & Capacity

Ini bagian penting yang sering bikin bingung pemula. Slice punya dua "ukuran":

- **Length (panjang)** → jumlah elemen yang **sedang ada** di slice
- **Capacity (kapasitas)** → jumlah elemen maksimal yang bisa ditampung **sebelum slice perlu memperbesar dirinya sendiri**

```go
s := make([]int, 3, 5) // length = 3, capacity = 5
fmt.Println(len(s)) // 3
fmt.Println(cap(s)) // 5
```

Bayangkan capacity seperti ukuran kardus, dan length seperti jumlah barang yang sudah dimasukkan ke kardus itu. Kardus bisa lebih besar dari isinya saat ini.

---

## 4. Operasi Dasar Slice

### Menambah elemen — `append()`

```go
buah := []string{"Apel", "Jeruk"}
buah = append(buah, "Mangga")
fmt.Println(buah) // [Apel Jeruk Mangga]
```

⚠️ **Penting:** `append()` **tidak mengubah slice secara langsung** di beberapa kasus, makanya hasilnya harus ditampung ulang ke variabel (`buah = append(buah, ...)`).

### Mengakses elemen (index dimulai dari 0)

```go
fmt.Println(buah[0]) // Apel
```

### Mengambil sebagian slice (slicing)

```go
angka := []int{10, 20, 30, 40, 50}
potongan := angka[1:3] // ambil index 1 sampai sebelum index 3
fmt.Println(potongan) // [20 30]
```

Format umum: `slice[mulai:selesai]` — index `selesai` **tidak ikut termasuk**.

### Menghapus elemen (Go tidak punya fungsi bawaan, jadi pakai trik ini)

```go
angka := []int{10, 20, 30, 40, 50}
// menghapus index ke-2 (nilai 30)
angka = append(angka[:2], angka[3:]...)
fmt.Println(angka) // [10 20 40 50]
```

### Looping slice

```go
for i, nilai := range angka {
    fmt.Println(i, nilai)
}
```

---

## 5. Hal Penting yang Sering Bikin Bingung Pemula

### Slice itu "menunjuk" ke array yang sama (reference)

Kalau kamu buat slice baru dari slice lama tanpa `append` yang menambah kapasitas, **keduanya bisa saling memengaruhi** karena sama-sama menunjuk ke data yang sama di memori.

```go
a := []int{1, 2, 3}
b := a
b[0] = 99
fmt.Println(a) // [99 2 3] — ikut berubah!
```

Ini beda dengan array biasa, yang kalau di-copy, benar-benar jadi data baru yang terpisah.

---

## 6. Studi Kasus Sederhana: Aplikasi Daftar Belanja

Supaya lebih kebayang, kita bikin studi kasus kecil: sistem **daftar belanja** sederhana menggunakan slice.

### Kasus:

Seorang pengguna ingin mencatat daftar belanjaan, menambah barang baru, menghapus barang yang sudah dibeli, dan menampilkan semua barang yang tersisa.

```go
package main

import "fmt"

func main() {
    // 1. Membuat daftar belanja awal
    daftarBelanja := []string{"Beras", "Telur", "Gula"}
    fmt.Println("Daftar awal:", daftarBelanja)

    // 2. Menambah barang baru ke daftar
    daftarBelanja = append(daftarBelanja, "Minyak Goreng")
    fmt.Println("Setelah ditambah:", daftarBelanja)

    // 3. Barang "Telur" sudah dibeli, jadi dihapus dari daftar
    index := 1 // index dari "Telur"
    daftarBelanja = append(daftarBelanja[:index], daftarBelanja[index+1:]...)
    fmt.Println("Setelah dihapus:", daftarBelanja)

    // 4. Menampilkan semua barang yang tersisa dengan nomor urut
    fmt.Println("\nSisa belanjaan yang perlu dibeli:")
    for i, barang := range daftarBelanja {
        fmt.Printf("%d. %s\n", i+1, barang)
    }
}
```

### Output yang dihasilkan:

```
Daftar awal: [Beras Telur Gula]
Setelah ditambah: [Beras Telur Gula Minyak Goreng]
Setelah dihapus: [Beras Gula Minyak Goreng]

Sisa belanjaan yang perlu dibeli:
1. Beras
2. Gula
3. Minyak Goreng
```

### Penjelasan studi kasus:

| Langkah             | Fungsi Slice yang Dipakai           | Kegunaan                                 |
| ------------------- | ----------------------------------- | ---------------------------------------- |
| Membuat daftar awal | Slice literal `[]string{...}`       | Menyimpan data awal                      |
| Menambah barang     | `append()`                          | Menambahkan elemen baru di akhir slice   |
| Menghapus barang    | `append(slice[:i], slice[i+1:]...)` | Menghilangkan elemen berdasarkan index   |
| Menampilkan data    | `range`                             | Melakukan iterasi (looping) semua elemen |

Studi kasus ini menunjukkan kenapa slice sangat cocok dipakai untuk data yang **jumlahnya tidak pasti** dan **sering berubah**, seperti daftar belanja, daftar tugas (to-do list), atau daftar pengguna aktif dalam sebuah aplikasi.

---

## 7. Ringkasan

| Konsep       | Penjelasan Singkat                                     |
| ------------ | ------------------------------------------------------ |
| Slice        | Struktur data dinamis, dibangun di atas array          |
| `len()`      | Jumlah elemen saat ini                                 |
| `cap()`      | Kapasitas maksimum sebelum slice membesar otomatis     |
| `append()`   | Menambah elemen ke slice                               |
| `slice[a:b]` | Mengambil sebagian slice dari index a sampai sebelum b |
| Reference    | Slice yang di-copy tetap menunjuk ke data yang sama    |

---

## 8. Latihan Mandiri (Opsional)

Coba kembangkan sendiri studi kasus di atas dengan:

1. Menambahkan fungsi untuk mengecek apakah suatu barang sudah ada di daftar sebelum ditambahkan (hindari duplikat).
2. Membuat fungsi terpisah `TambahBarang()` dan `HapusBarang()` agar kode lebih rapi.
3. Menyimpan slice barang beserta jumlah/quantity-nya menggunakan `struct`.

Latihan ini akan membantu memahami slice lebih dalam sekaligus mengenalkan konsep function dan struct di Go.
