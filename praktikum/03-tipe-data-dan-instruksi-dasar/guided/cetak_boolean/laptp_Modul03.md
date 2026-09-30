# <h1 align="center">Tugas Pendahuluan Modul 003 - Tipe Data dan Instruksi Dasar</h1>
<p align="center">Nur Widodo - 109092630005O</p>

### 1. cetak_boolean.go

```go
package main

import "fmt"

func main() {
	var bool bool

	fmt.Println("Masukkan nila boolean")
	fmt.Scan(&bool)

	fmt.Println(bool)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](..\cetak_boolean\cetak_boolean.png)


#### Deskripsi
Program membaca sebuah nilai boolean (true atau false) dari input pengguna, lalu mencetaknya kembali apa adanya. Nilai disimpan dalam variabel bool bertipe bool.

Hasil uji:
- input true → keluaran true
- input false → keluaran false

## Kesimpulan
Program ini mempraktikkan cara membaca dan mencetak nilai boolean di Go. Nilai true atau false disimpan dalam variabel bertipe bool, dibaca dengan fmt.Scan, lalu dicetak kembali apa adanya. Hasil pengujian menunjukkan input true menghasilkan keluaran true dan input false menghasilkan keluaran false, sesuai dengan yang diminta soal.