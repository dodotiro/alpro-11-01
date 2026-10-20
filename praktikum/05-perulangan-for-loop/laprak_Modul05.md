# <h1 align="center">Laporan Praktikum Modul 05 - Perulangan For</h1>
<p align="center">Nur Widodo - 109092630005</p>

## Dasar Teori

### A. Pengenalan for
Perulangan for merupakan looping construct satu-satunya dalam Go, karena Go tidak memiliki while dan do-while. Fungsi utama dari for yaitu untuk mengeksekusi kode didalamnya secara berulang selama kondisi-kondisinya terpenuhi.
### B. Cara Penggunaan for di Go
Berikut adalah keterangan komponen utama pada perulangan standar for di Go beserta penjelasannya untuk laporan praktikum:

```
for i := 1; i <= n; i++ {
    // blok kode
}

```

Keterangan Komponen:
1. Inisialisasi (i := 1): 
   Mendeklarasikan dan memberi nilai awal pada variabel penghitung (counter) sebelum perulangan dimulai. Variabel ini hanya berlaku di dalam blok for.
2. Kondisi (i <= n): 
   Ekspresi boolean yang dievaluasi sebelum setiap iterasi. Jika bernilai true, blok kode di dalam for akan dijalankan. Jika false, perulangan berhenti.
3. Post-statement (i++): 
   Perintah yang dieksekusi di akhir setiap iterasi (biasanya menaikkan atau menurunkan nilai counter).

## Guided

### 1. tabungan.go

```go
package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Scan(&a, &b, &n)

	total := 0

	TabunganHariIni := a

	for i := 1; i <= n; i++ {
		total = total + TabunganHariIni
		TabunganHariIni = TabunganHariIni + b

	}

	fmt.Println(total)
}
```

#### Deskripsi
Program ini menghitung total tabungan. Dimana a sebagai nominal tabungan hari pertama dan b menambah tabungan setiap harinya dalam bentuk rupiah.

### 2. terbesar.go

```go
package main

import "fmt"

func main() {
	var n, angka, max int
	fmt.Scan(&n, &angka)

	max = angka
	for i := 2; i <= n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}
	}
	fmt.Println(max)
}

```

#### Deskripsi
Program ini mencari bilangan terbesar dari n buah bilangan bulat yang dimasukkan.

### 3. bioskop.go

```go
package main

import "fmt"

func main() {
	var t int
	fmt.Print("Input t (jumlah transaksi): ")
	fmt.Scan(&t)
	totalPendapatan := 0

	for i := 1; i <= t; i++ {
		var jenisFilm, hari string
		var jumlahTiket int

		fmt.Printf("Input transaksi ke-%d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

		hargaSatuan := 0

		if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			if jenisFilm == "Reguler" {
				hargaSatuan = 40000
			} else if jenisFilm == "IMAX" {
				hargaSatuan = 80000
			} else if jenisFilm == "3D" {
				hargaSatuan = 60000
			}
		} else if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			if jenisFilm == "Reguler" {
				hargaSatuan = 50000
			} else if jenisFilm == "IMAX" {
				hargaSatuan = 90000
			} else if jenisFilm == "3D" {
				hargaSatuan = 70000
			}
		}
		totalHarga := hargaSatuan * jumlahTiket

		if jumlahTiket >= 5 {
			diskon := (totalHarga * 10) / 100
			totalHarga = totalHarga - diskon
		}
		fmt.Printf("Output total harga transaksi ke-%d: %d\n", i, totalHarga)

		totalPendapatan = totalPendapatan + totalHarga
	}
	fmt.Println("Output total pendapatan: ", totalPendapatan)
}

```

#### Deskripsi
Program ini menghitung total harga tiket bioskop dimana jika beli >= 5 tiket maka mendapatkan diskon 10%.

## Unguided

### 1. prima.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			count++
		}
	}
	fmt.Println()

	if count == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}

```

##### Output
![Screenshot Output Unguided](./Unguided/prima/output.png)

#### Deskripsi
Program ini menerima sebuah bilangan bulat positif n, kemudian mencari dan menampilkan seluruh faktor dari bilangan tersebut menggunakan perulangan for. Setelah itu, program menghitung jumlah faktornya. Jika jumlah faktor tepat dua, yaitu 1 dan bilangan itu sendiri, program menampilkan “Ya” yang berarti bilangan tersebut prima. Jika tidak, program menampilkan “Tidak”.

### 2. ecct.go

```go
package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)

	totalLulus := 0

	for i := 0; i < m; i++ {
		var p, v, ga, fs, ci, tc float64
		fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

		rataRata := (p + v + ga + fs + ci + tc) / 6.0

		if rataRata >= 3.33 {
			fmt.Println("Mendapat Sertifikat")
			totalLulus++
		} else {
			fmt.Println("Tidak Mendapat Sertifikat")
		}
	}

	fmt.Println(totalLulus)
}

```

##### Output
![Screenshot Output Unguided](./Unguided/ecct/output.png)

#### Deskripsi
Program menghitung rata-rata nilai menggunakan perulangan for. Mahasiswa mendapat sertifikat jika rata-ratanya minimal 3,33. Setelah memeriksa seluruh mahasiswa, program menampilkan status masing-masing mahasiswa serta jumlah mahasiswa yang mendapat sertifikat.

## Kesimpulan
Semua program diatas menggunakan fungsi perulangan for


## Referensi
1. Go Team. (2026). _The Go Programming Language Specification_. Google LLC. Diakses pada 04 Oktober 2026 melalui https://go.dev/ref/spec.
