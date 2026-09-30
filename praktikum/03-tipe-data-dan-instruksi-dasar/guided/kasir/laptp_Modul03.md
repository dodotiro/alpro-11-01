# <h1 align="center">Tugas Pendahuluan Modul 003 - Tipe Data dan Instruksi Dasar</h1>
<p align="center">Nur Widodo - 109092630005O</p>

### 1. kasir.go

```go
package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var sepuluhRibuan int = x / 10000
	var sisa int = x % 10000

	var limaRibuan int = sisa / 5000
	sisa = sisa % 5000

	var seribuan int = sisa / 1000

	fmt.Println(sepuluhRibuan, limaRibuan, seribuan)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](../kasir/kasir.png)

#### Deskripsi
Program ini ditulis dalam bahasa Go (Golang) untuk memecah atau mengonversi suatu jumlah uang (bilangan bulat x) ke dalam pecahan lembar uang yang lebih kecil, yaitu pecahan 10.000, 5.000, dan 1.000. Program menggunakan operasi pembagian bulat (/) untuk menentukan jumlah lembar uang dan operator sisa bagi atau modulo (%) untuk membawa sisa nilai uang ke pecahan berikutnya secara berurutan.

## Kesimpulan
- Pecahan Nilai Beruntun: Memanfaatkan kombinasi operator pembagian (/) dan modulo (%) secara berantai untuk menghitung jumlah lembar pecahan uang dari nominal terbesar hingga terkecil berdasarkan sisa nominal sebelumnya.
- Efisiensi Alur Variabel: Menerapkan pembaruan nilai variabel sisa secara fleksibel (sisa = sisa % 5000) untuk melacak sisa uang yang belum dikonversi pada setiap tahapan pecahan.