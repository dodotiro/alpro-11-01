# <h1 align="center">Laporan Praktikum Modul 02 - Algoritma dan Pemrograman Variabel, Tipe Data, dan Operasi</h1>
<p align="center">Nur Widodo - 109092630005O</p>

## Dasar Teori

### A. Bahasa Pemrograman Go
Menurut Donovan & Kernighan (2015), Go (atau Golang) adalah bahasa pemrograman open-source yang dirancang untuk memudahkan pembangunan perangkat lunak yang sederhana, cepat, dan handal. Go menyediakan dukungan bawaan untuk konkurensi serta sistem tipe data yang statis namun fleksibel, sehingga sangat cocok digunakan untuk pengembangan aplikasi sistem maupun komputasi modern.

### B. Variabel, Tipe Data, dan Operator Aritmatika di Go

#### 1. Deklarasi Variabel dan Tipe Data
Dalam bahasa pemrograman Go, variabel digunakan untuk menyimpan nilai data yang dapat diubah selama eksekusi program. Pada program pemecahan pecahan uang ini, tipe data int digunakan secara menyeluruh untuk merepresentasikan nominal uang, hasil pembagian pecahan, serta sisa perhitungan modulus

#### 2. Operator Aritmatika (Pembagian dan Modulus)
Operator pembagian (/) dan modulus (%) memainkan peran krusial dalam pemecahan masalah konversi nilai nominal uang. Operator pembagian bilangan bulat menghasilkan jumlah lembar atau keping pecahan tertentu, sementara operator modulus digunakan untuk menghitung sisa nominal uang yang belum dikonversi ke pecahan yang lebih kecil

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Unguided

### 1. cacahuang.go

```go
package main

import "fmt"

func main() {
	var uang int
	fmt.Scan(&uang)

	sepuluhRibu := uang / 10000
	sisa := uang % 10000

	limaRibu := sisa / 5000
	sisa = sisa % 5000

	seribu := sisa / 1000

	total := sepuluhRibu + limaRibu + seribu
	fmt.Println(total)
	fmt.Println(sepuluhRibu, limaRibu, seribu)

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Preview Output](/cacahuang/cacahuang.png)


#### Deskripsi
Praktikum ini membahas pembuatan program untuk menghitung pecahan mata uang berdasarkan total nominal uang yang diinputkan pengguna. Program menggunakan operator aritmatika pembagian bulat (/) dan sisa bagi atau modulus (%) untuk memecah total uang ke dalam pecahan nominal Rp10.000, Rp5.000, dan Rp1.000. Program kemudian menampilkan total lembar pecahan serta rincian jumlah lembar untuk masing-masing pecahan tersebut.


## Kesimpulan
Berdasarkan praktikum yang telah dilakukan, dapat disimpulkan bahwa:

1. Bahasa pemrograman Go sangat efisien dalam memproses perhitungan matematis menggunakan tipe data int dan operator dasar.
2. SKombinasi operator pembagian (/) dan modulus (%) sangat efektif untuk menyelesaikan masalah komputasi terstruktur seperti konversi pecahan nilai mata uang.
3. Alur logika program yang runtut dari pembacaan input, proses kalkulasi bertahap, hingga pencetakan hasil akhir memastikan program berjalan dengan akurat dan sesuai harapan.


## Referensi
1. Sumber resmi Golang https://go.dev/
2. Jurnal modul 02 - Golang - Telkom University

<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
