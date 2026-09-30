# <h1 align="center">Tugas Pendahuluan Modul 003 - Tipe Data dan Instruksi Dasar</h1>
<p align="center">Nur Widodo - 109092630005O</p>

### 1. konversi_suhu.go

```go
package main

import "fmt"

func main() {
	var c float64

	fmt.Println("Masukkan suhu celcius")
	fmt.Scan(&c)

	r := (4.0 / 5.0) * c
	fmt.Println(r, "Reamur")

}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](../konversi_suhu/konversi_suhu.png)


#### Deskripsi
Program ini ditulis dalam bahasa Go (Golang) untuk mengonversi suhu dari derajat Celsius menjadi derajat Reamur. Program menerima masukan bilangan real (floating-point) yang merepresentasikan suhu dalam Celsius, kemudian menghitung dan menampilkan hasilnya menggunakan rumus konversi:

$$R = (4/5) \times C$$

## Kesimpulan
- Penggunaan Tipe Data Pecahan: Memanfaatkan tipe data float64 untuk menangani bilangan real agar hasil perhitungan suhu presisi dan akurat.
- Penerapan Rumus Matematis: Menerapkan operasi aritmatika dasar Go dengan membedakan bentuk pecahan (4.0 / 5.0) guna mencegah pembagian bulat (integer division).