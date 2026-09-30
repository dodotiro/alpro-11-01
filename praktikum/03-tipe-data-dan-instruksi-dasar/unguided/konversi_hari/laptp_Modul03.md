# <h1 align="center">Tugas Pendahuluan Modul 003 - Tipe Data dan Instruksi Dasar</h1>
<p align="center">Nur Widodo - 109092630005O</p>

### 1. konversi_hari.go

```go
package main

import "fmt"

func main() {
	var totalHari int
	fmt.Print("Masukkan Jumlah Hari: ")
	fmt.Scan(&totalHari)

	hariPertahun := 360
	hariPerbulan := 30
	hariPerminggu := 7

	tahun := totalHari / hariPertahun
	sisaHari := totalHari % hariPertahun

	bulan := sisaHari / hariPerbulan
	sisaHari = sisaHari % hariPerbulan

	minggu := sisaHari / hariPerminggu
	hari := sisaHari % hariPerminggu

	fmt.Printf("%d hari setara dengan: %d tahun, %d bulan, %d minggu, dan %d hari.\n", totalHari, tahun, bulan, minggu, hari)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](../konversi_hari/konversi_hari.png)


#### Deskripsi
Program ini ditulis dalam bahasa Go (Golang) untuk mengonversi jumlah hari ke dalam satuan waktu yang lebih besar, yaitu tahun, bulan, minggu, dan sisa hari.

Program dirancang menggunakan aturan perhitungan waktu standar konversi sederhana:

- 1 tahun = 12 bulan (360 hari)
- 1 bulan = 30 hari
- 1 minggu = 7 hari

Melalui pendekatan operasi pembagian bulat (/) dan sisa bagi atau modulo (%), program secara beruntun memecah total hari masukan dari pengguna hingga mendapatkan sisa hari terkecil.

## Kesimpulan
- Logika Konversi Beruntun: Memanfaatkan operator pembagian (/) dan modulo (%) secara berantai di mana satuan waktu berikutnya dihitung dari sisa hari sebelumnya.
- Aturan Sintaks Go: Mematuhi aturan wajib deklarasi variabel, memastikan setiap variabel digunakan, serta menggunakan fmt.Printf untuk pencetakan berformat.