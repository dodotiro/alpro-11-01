# <h1 align="center">Tugas Pendahuluan Modul 003 - Tipe Data dan Instruksi Dasar</h1>
<p align="center">Nur Widodo - 109092630005O</p>

### 1. konversi.go

```go
package main

import "fmt"

func main() {
	var mil float64

	fmt.Println("Masukkan jarak mil")
	fmt.Scan(&mil)

	fmilkm := mil * 1.6
	fmt.Println(fmilkm, "km")

}


```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](../konversi_mil_ke_mil/konversi.png)

#### Deskripsi
Program membaca sebuah nilai boolean (true atau false) dari input pengguna, lalu mencetaknya kembali apa adanya. Nilai disimpan dalam variabel bool bertipe bool.

Hasil uji:
- input true → keluaran true
- input false → keluaran false


## Kesimpulan
Program ini mempraktikkan konversi satuan jarak menggunakan tipe data float64 dan perkalian. Nilai jarak dalam mil dikalikan dengan 1.6 untuk menghasilkan nilai dalam kilometer. Dari pengujian, input 1 menghasilkan 1.6 km. Program juga mempraktikkan cara memformat bilangan desimal: dengan fmt.Printf("%.1fkm", hasil) keluaran menjadi tepat satu angka di belakang koma, sehingga input 1.25 menghasilkan 2.0km dan input 2 menghasilkan 3.2km.