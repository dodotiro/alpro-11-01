# <h1 align="center">Laporan Praktikum Modul 02 - Algoritma dan Pemrograman Variabel, Tipe Data, dan Operasi</h1>
<p align="center">Nur Widodo - 109092630005O</p>

## Dasar Teori

### A. Bahasa Pemrograman Go
Menurut Donovan & Kernighan (2015), Go (atau Golang) adalah bahasa pemrograman open-source yang dirancang untuk memudahkan pembangunan perangkat lunak yang sederhana, cepat, dan handal. Go menyediakan dukungan bawaan untuk konkurensi serta sistem tipe data yang statis namun fleksibel, sehingga sangat cocok digunakan untuk pengembangan aplikasi sistem maupun komputasi modern.

### B. Variabel, Tipe Data, dan Operator Aritmatika di Go

#### 1. Deklarasi Variabel dan Tipe Data
Dalam bahasa pemrograman Go, variabel dideklarasikan dengan tipe data tertentu secara eksplisit maupun implisit. Pada program kalkulator ini, digunakan tipe data int untuk merepresentasikan bilangan bulat pada variabel angka pertama dan kedua, serta tipe data string untuk menyimpan operator aritmatika (+, -, *, /, %)

#### 2. Struktur Kontrol Percabangan dan Operator
Percabangan switch digunakan di Go untuk mengevaluasi ekspresi operator yang dimasukkan pengguna dan mengarahkan eksekusi ke blok kode yang sesuai (penjumlahan, pengurangan, perkalian, pembagian, atau sisa bagi/modulus). Validasi kondisi juga disematkan untuk mencegah kesalahan pembagian atau modulus dengan angka nol (b == 0), guna menghindari error runtime pada program.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Unguided

### 1. kalkulator.go

```go
package main

import "fmt"

func main() {
	var a, b int
	var operator string

	//Input angka pertama
	fmt.Print("Masukkan Angka Pertama :")
	fmt.Scan(&a)

	//Pilih operator
	fmt.Print("Masukkan Operator (+, -, *, /, %): ")
	fmt.Scan(&operator)

	//Input angka kedua
	fmt.Print("Masukkan Angka Kedua :")
	fmt.Scan(&b)

	//Menampilkan hasil
	switch operator {
	case "+":
		fmt.Printf("Hasil: %d + %d = %d\n", a, b, a+b)
	case "-":
		fmt.Printf("Hasil: %d - %d = %d\n", a, b, a-b)
	case "*":
		fmt.Printf("Hasil: %d * %d = %d\n", a, b, a*b)
	case "/":
		if b == 0 {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan!")
		} else {
			fmt.Printf("Hasil: %d / %d = %d\n", a, b, a/b)
		}
	case "%":
		if b == 0 {
			fmt.Println("Error: Modulus dengan nol tidak diperbolehkan!")
		} else {
			fmt.Printf("Hasil: %d %% %d = %d\n", a, b, a%b) // Gunakan %% untuk mencetak simbol % di Printf
		}
	default:
		fmt.Println("Operator tidak dikenal!")
	}
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
..\alpro-11-01\praktikum\02-bahasa-pemrograman-go\unguided\kalkulator\kalkulator.png


#### Deskripsi
Praktikum ini membahas pembuatan program kalkulator sederhana menggunakan bahasa Go. Program menerima tiga input dari pengguna, yaitu angka pertama, operator aritmatika, dan angka kedua. Berdasarkan operator yang dipilih, program menggunakan struktur kontrol switch-case untuk mengeksekusi operasi aritmatika yang sesuai (+, -, *, /, dan %). Selain itu, program juga dilengkapi dengan validasi logika kondisi guna menangani error apabila pengguna mencoba melakukan pembagian atau sisa modulus dengan angka nol.


## Kesimpulan
Berdasarkan praktikum yang telah dilakukan, dapat disimpulkan bahwa:

1. Bahasa pemrograman Go menyediakan tipe data dasar seperti int dan string serta fungsi input-output melalui package fmt yang sangat mendukung pembuatan program interaktif sederhana.
2. Struktur percabangan switch-case sangat efektif digunakan untuk menangani percabangan kondisi jamak seperti pemilihan operator kalkulator.
3. Penanganan kondisi khusus (seperti validasi pembagian dengan nol) sangat penting untuk mencegah gangguan eksekusi program atau runtime error.


## Referensi
1. Sumber resmi Golang https://go.dev/
2. Jurnal modul 02 - Golang - Telkom University

<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
