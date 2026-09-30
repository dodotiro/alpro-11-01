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

	k := c + 273
	fmt.Println(k, "kelvin")

}
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](../konversi_suhu/konversi_suhu.png)


#### Deskripsi
Program ini ditulis dalam bahasa Go (Golang) untuk mengonversi suhu dari derajat Celsius menjadi derajat Kelvin. Program menerima masukan bilangan real (floating-point) yang merepresentasikan suhu dalam Celsius, kemudian menghitung dan menampilkan hasilnya menggunakan rumus Kelvin:

$$K = C \times 273$$

## Kesimpulan
- Operasi Aritmatika Sederhana: Memanfaatkan penambahan dasar (c + 273) untuk mengonversi nilai suhu dari skala Celsius ke Kelvin.
- Penggunaan Tipe Data Bilangan Real: Menggunakan tipe data float64 (var c float64) untuk menampung nilai masukan suhu agar tetap fleksibel dan akurat.