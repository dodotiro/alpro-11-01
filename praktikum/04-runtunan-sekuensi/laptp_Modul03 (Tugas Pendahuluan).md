# <h1 align="center">Tugas Pendahuluan Modul 004 - Runtunan Sekuensi</h1>
<p align="center">Nur Widodo - 109092630005</p>

### 1. evaluasi_expresi_kontrol.go

```go
package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	if sngNum > 0 || (intNum >= 0 && -1*intOther == -10) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP/evaluasi_ekpresi_kontrol/output.png)

#### Deskripsi
Program yang berfungsi memeriksa ekspresi kontrol menghasilkan nilai "true" atau "false"

### 2. tracing.go

```go
package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}

	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}

	fmt.Println("Nilai akhir result:", result)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP/tracing/output.png)

#### Deskripsi
Program ini berfungsi untuk menentukan nilai akhir dari setiap variabel setelah program selesai dijalankan.

Pertanyaan:
1. Berapa nilai akhir dari variabel result setelah semua pernyataan kondisi dieksekusi?
result = 25
2. Apa output yang dihasilkan oleh program?
Nilai akhir result: 25
3. Tuliskan langkah-langkah alur eksekusi program berdasarkan kondisi yang diberikan:
	- Kondisi 1: x > 5 → true. Masuk blok, y < 10 juga true, jadi result = 10 + 5 = 15.

	- Kondisi 2: z > 10 && x == 10 → true dan true = true. result = 15 + 15 = 30.


	- Kondisi 3: x == 10 || y > 10 → true (cukup kiri yang true). result = 30 + 5 = 35.

### 3. jumlah_hari_dalam_tahun.go

```go
package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Println("Masukkan tahun: ")
	fmt.Scan(&tahun)
	fmt.Println("Masukkan bulan: ")
	fmt.Scan(&bulan)

	bulanKabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println("Jumlah hari: 31")
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println("Jumlah hari: 30")
	case "Feb":
		if bulanKabisat {
			fmt.Println("Jumlah hari: 29")
		} else {
			fmt.Println("Jumlah hari: 28")
		}
	default:
		fmt.Println("-")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP/jumlah_hari_dalam_tahun/output.png)

#### Deskripsi
Program ini menentukan jumlah hari pada setiap bulan dan tahun termasuk bulan kabisat februari.

Hasil uji:
- Februari non kabisat : 28
- Februari kabisat : 29

### 4. switch.go

```go
package main

import "fmt"

func main() {
	var os string

	fmt.Println("Masukkan Operasi Sistem: ")
	fmt.Scan(&os)

	switch os {
	case "windows", "macos":
		fmt.Println("Close Source")
	case "linux":
		fmt.Println("Open Source")
	default:
		fmt.Println("-")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP/switch_case/output.png)

#### Deskripsi
Program ini menerapkan penggunaan "switch case" dan dalam contohnya untuk menentukan sistem operasi (OS) Open source atau Close source.

## Kesimpulan
Praktikum ini mempraktikkan penggunaan switch, case, dan default dalam sebuah program

Program pertama memeriksa ekspresi kontrol menghasilkan nilai "true" atau "false"

Program kedua mberfungsi untuk menentukan nilai akhir dari setiap variabel setelah program selesai dijalankan.

Program ketiga menentukan jumlah hari pada setiap bulan dan tahun termasuk bulan kabisat februari.

Hasil uji:
- Februari non kabisat : 28
- Februari kabisat : 29

Program keempat menerapkan penggunaan "switch case" dan dalam contohnya untuk menentukan sistem operasi (OS) Open source atau Close source.

Secara keseluruhan, ketiga program berjalan sesuai dengan yang diminta soal dan menunjukkan pemahaman dalam pemilihan tipe data yang tepat untuk setiap kebutuhan, penggunaan operator aritmatika, serta pembacaan dan pencetakan nilai dari input pengguna.
