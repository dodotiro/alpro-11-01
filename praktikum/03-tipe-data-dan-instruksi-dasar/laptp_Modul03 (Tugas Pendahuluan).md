# <h1 align="center">Tugas Pendahuluan Modul 003 - Tipe Data dan Instruksi Dasar</h1>
<p align="center">Nur Widodo - 109092630005</p>

### 1. sisa_kue.go

```go
package main

import "fmt"

func main() {
	var x, y int

	fmt.Println("Masukkan jumlah kue: ")
	fmt.Scan(&y)

	fmt.Println("Masukkan jumlah anggota keluarga: ")
	fmt.Scan(&x)

	sisa := y % x
	fmt.Println("Sisa kue adalah: ", sisa)

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/sisa_kue/sisa_kue.png)

#### Deskripsi
Program membaca dua bilangan bulat, yaitu y (jumlah kue) dan x (jumlah anggota keluarga). Jumlah kue sisa setelah dibagi rata dihitung menggunakan operator sisa bagi (%), yang menghasilkan sisa pembagian y dibagi x. Nilai sisa tersebut kemudian dicetak.

### 2. konversi.go

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
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/konversi_mil_ke_km/konversi.png)

#### Deskripsi
Program membaca jarak dalam satuan mil dari input pengguna, lalu mengonversinya ke kilometer dengan mengalikan nilai tersebut dengan 1.6. Hasil konversi disimpan dalam variabel fmilkm bertipe float64 dan dicetak ke layar.

Hasil uji:
- input 1 → keluaran 1.6 km
- input 1.25 → keluaran 2 km
- input 2 → keluaran 3.2 km

### 3. cetak_boolean.go

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
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/cetak_boolean/cetak_boolean.png)

#### Deskripsi
Program membaca sebuah nilai boolean (true atau false) dari input pengguna, lalu mencetaknya kembali apa adanya. Nilai disimpan dalam variabel bool bertipe bool.

Hasil uji:
- input true → keluaran true
- input false → keluaran false

## Kesimpulan
Praktikum ini mempraktikkan penggunaan tipe data dasar dan instruksi dasar pada bahasa Go, meliputi tipe int, float64, dan bool beserta proses input-output dan operasi aritmatika sederhana.

Program pertama mempraktikkan operator sisa bagi (%) pada bilangan bulat: jumlah kue y dibagi rata kepada x anggota keluarga dan sisanya dicetak. Dari pengujian, 10 kue untuk 5 orang menghasilkan sisa 0, sedangkan 11 kue untuk 5 orang menghasilkan sisa 1. Operator % hanya berlaku pada tipe bilangan bulat, sehingga variabelnya harus bertipe int.

Program kedua mempraktikkan konversi satuan jarak menggunakan tipe float64 dan perkalian: jarak dalam mil dikalikan 1.6 menghasilkan jarak dalam kilometer. Dari pengujian, input 1 menghasilkan 1.6 km dan input 2 menghasilkan 3.2 km. Program ini juga menunjukkan cara memformat bilangan desimal, misalnya dengan fmt.Printf("%.1fkm", hasil) agar keluaran tepat satu angka di belakang koma.

Program ketiga mempraktikkan tipe bool: nilai true atau false dibaca dengan fmt.Scan lalu dicetak kembali apa adanya, dan pengujian menunjukkan keluaran selalu sama dengan masukan.

Secara keseluruhan, ketiga program berjalan sesuai dengan yang diminta soal dan menunjukkan pemahaman dalam pemilihan tipe data yang tepat untuk setiap kebutuhan, penggunaan operator aritmatika, serta pembacaan dan pencetakan nilai dari input pengguna.
