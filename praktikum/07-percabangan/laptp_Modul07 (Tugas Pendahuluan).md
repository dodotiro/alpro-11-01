# <h1 align="center">Tugas Pendahuluan Modul 07 - Percabangan</h1>
<p align="center">Nur Widodo - 109092630005</p>

### 1. Perbedaan switch pada Go & Java
1. Percabangan di Go

* **`if`**: Cek 1 kondisi benar.
* **`if-else`**: Cek kondisi benar atau salah.
* **`else if`**: Cek banyak kondisi secara berurutan.
* **`switch`**: Pengganti `else if` untuk mencocokkan banyak pilihan.

2. `switch` Go vs Java

1. **Tanpa `break**`: Go otomatis berhenti di tiap `case` (Java butuh `break`).
2. **Tanpa Ekspresi**: Go bisa membuat `switch` tanpa variabel utama, isi `case` bebas berupa kondisi boolean.

### 2. bilangan.go.go

```go
package main
import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x == 0 {
		fmt.Println("Nol")
	} else if x % 2 == 0 {
		fmt.Println("Genap")
		if x < 0 {
			fmt.Println("Negatif")
		} else if x >= 0 {
			fmt.Println("Positif")
		}
	} else if x % 2 != 0 {
		fmt.Println("Ganjil")
		if x < 0 {
			fmt.Print("Negatif")
		} else if x >= 0 {
			fmt.Println("Positif")
		}
	}
}
```

#### Output
![Screenshot Output Unguided](./tp/bilangan/output.png)


#### Deskripsi
Program ini menentukan jenis bilangan apakah bilangan tersebut merupakan bilangan ganjil / genap / nol.

### 3. hari.go

```go
package main

import "fmt"

func main() {
	var hari int
	fmt.Scan(&hari)

	switch {
	case hari == 1:
		fmt.Println("Senin")
	case hari == 2:
		fmt.Println("Selasa")
	case hari == 3:
		fmt.Println("Rabu")
	case hari == 4:
		fmt.Println("Kamis")
	case hari == 5:
		fmt.Println("Jumat")
	case hari == 6:
		fmt.Println("Sabtu", "Akhir Pekan")
	case hari == 7:
		fmt.Println("Minggu", "Akhir Pekan")
	default:
		fmt.Println("Nomor hari tidak valid")
	}
}

```

##### Output
![Screenshot Output Unguided](./tp/hari/output.png)


#### Deskripsi
Program ini berfungsi menentukan urutan hari dimulai dari hari 1 yaitu senin dan hari 7 yaitu minggu. dan jika hari > 7 maka hasilnya tidak valid.

### 4. terbesar.go

```go
package main

import "fmt"

func main() {
	var a, b, c, terbesar int
	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		terbesar = a
	} else if b >= a && b >= c {
		terbesar = b
	} else {
		terbesar = c
	}
	fmt.Println(terbesar)
}

```
#### Output
![Screenshot Output Unguided](./tp/terbesar/output.png)


#### Deskripsi
Program ini menentukan sebuah bilangan terbesar dari beberapa urutan bilangan yang ada.

## Kesimpulan
Semua program di atas menggunakan Percabangan dan fungsi if else switch case untuk membuat sebuah program.