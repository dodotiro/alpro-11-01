# <h1 align="center">Laporan Praktikum Modul 07 - Percabangan</h1>
<p align="center">Nur Widodo - 109092630005</p>

## Dasar Teori

### A. Pengenalan Percabangan
Dalam bahasa pemrograman Go (Golang), percabangan digunakan untuk mengambil keputusan berdasarkan kondisi tertentu. Berbeda dengan beberapa bahasa pemrograman lain, Go memiliki beberapa ciri khas dalam penulisan percabangannya (seperti tidak memerlukan tanda kurung () pada kondisi).

## Guided

### 1. kabisat.go

```go
package main

import "fmt"

func main() {
	var tahun int

	fmt.Scan(&tahun)

	if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
		println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}

```

#### Deskripsi
Program ini untuk menentukan tahun kabisat.

### 2. segitiga.go

```go
package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Bukan Segitiga")
	} else if a == b && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if a == b || a == c || b == c {
		println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}

```

#### Deskripsi
Program ini mencari jenis segitiga sama sisi, sama kaki, atau sembarang dari validasi panjang sisi.

### 3. kalkulator.go

```go
package main

import "fmt"

func main() {
	var a, b int
	var op string

	fmt.Scan(&a, &op, &b)

	switch op {
	case "+":
		fmt.Print(a + b)
	case "-":
		fmt.Print(a - b)
	case "*":
		fmt.Print(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Error: Pembagian dengan nol")
		} else {
			fmt.Printf("%.2f\n", float64(a)/float64(b))
		}
	case "%":
		if b == 0 {
			fmt.Println("Error: Pembagian dengan nol")
		} else {
			fmt.Println(a % b)
		}
	default:
		fmt.Println("Operator tidak valid")
	}
}

```

#### Deskripsi
Program ini merupakan kalkulator sederhana yang memiliki operator penjumlahan (+), perkalian (*), pembagian (/), dan modulo (%).

### 4. ongkir.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var berat float64
	var zona int
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Println("Berat melebihi batas maksimum (30kg)")
		return
	}
	var tarifPerKg int

	switch zona {
	case 1:
		tarifPerKg = 10000
	case 2:
		tarifPerKg = 15000
	case 3:
		tarifPerKg = 25000
	default:
		fmt.Println("Zona tidak valid")
		return
	}

	beratBulat := int(math.Ceil(berat))
	total := beratBulat * tarifPerKg

	if layanan == "Express" {
		total += total / 2
	}

	fmt.Println(beratBulat)
	fmt.Println(total)
}

```

#### Deskripsi
Program ini berfungsi menghitung ongkir (ongkos kirim) sesuai dengan berat, zona, dan layanan yang di pilih.

## Unguided

### 1. besok.go

```go
package main

import "fmt"

func main() {
	var tanggal, bulan, tahun int

	fmt.Scan(&tanggal, &bulan, &tahun)

	kabisat := tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0)

	jumlahHari := 0

	if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 ||
		bulan == 8 || bulan == 10 || bulan == 12 {
		jumlahHari = 31
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		jumlahHari = 30
	} else if bulan == 2 {
		if kabisat {
			jumlahHari = 29
		} else {
			jumlahHari = 28
		}
	}

	if bulan < 1 || bulan > 12 || tanggal < 1 || tanggal > jumlahHari {
		fmt.Println("Tanggal tidak valid")
		return
	}

	fmt.Println("Valid")

	tanggal++

	if tanggal > jumlahHari {
		tanggal = 1
		bulan++

		if bulan > 12 {
			bulan = 1
			tahun++
		}
	}

	fmt.Printf("%d/%d/%d\n", tanggal, bulan, tahun)
}

```

##### Output
![Screenshot Output Unguided](./Unguided/besok/output.png)

#### Deskripsi
Program memeriksa validitas tanggal dan menentukan tanggal berikutnya. Jumlah hari ditentukan berdasarkan bulan dan tahun kabisat. Jika bulan berakhir, tanggal kembali ke 1 dan bulan atau tahun akan bertambah


### 2. bmi.go

```go
package main

import "fmt"

func main() {
	var berat, tinggiCm float64
	fmt.Scan(&berat, &tinggiCm)

	tinggiM := tinggiCm / 100.0
	bmi := berat / (tinggiM * tinggiM)

	fmt.Printf("%.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("Kurus")
	} else if bmi < 25 {
		fmt.Println("Normal")
	} else if bmi < 30 {
		fmt.Println("Gemuk")
	} else {
		fmt.Println("Obesitas")
	}
}
```

##### Output
![Screenshot Output Unguided](./Unguided/bmi/output.png)

#### Deskripsi
Program ini memeriksa BMI seseorang dari tinggi badan dan berat badan apakah normal atau obesitas.


## Kesimpulan
1. Percabangan dalam Go merupakan _control structure_ yang berfungsi seperti penunjuk arah, yaitu untuk menentukan jalur mana yang akan dieksekusi berdasarkan syarat tertentu.
2. Percabangan `if` menggunakan logika `true` atau `false` dalam menjalankan kode.
3. Percabangan `if` memiliki tiga jenis percabangan, yaitu `if`, `if-else`, dan `else-if`.
4. Percabangan `switch` mencocokkan nilai variabel dengan `case` yang ada didalamnya.
5. Percabangan `switch` akan otomatis berhenti ketika sudah ditemukan `case` yang memiliki nilai sesuai.


## Referensi
1. Go Team. (2026). _The Go Programming Language Specification_. Google LLC. Diakses pada 04 Oktober 2026 melalui https://go.dev/ref/spec.
