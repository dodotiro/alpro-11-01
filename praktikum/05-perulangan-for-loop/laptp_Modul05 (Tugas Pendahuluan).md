# <h1 align="center">Tugas Pendahuluan Modul 05 - Perulangan For</h1>
<p align="center">Nur Widodo - 109092630005</p>

### 1. Struktur Utama Perulangan For

Fungsi utama struktur perulangan for dalam Go adalah menjalankan blok kode secara berulang selama kondisi yang ditentukan masih bernilai true.

Tiga komponennya adalah:

1. Inisialisasi  
   Menentukan nilai awal variabel perulangan dan dijalankan satu kali sebelum perulangan dimulai.  
   Contoh: i := 1

2. Kondisi  
   Menentukan apakah perulangan masih boleh berjalan. Selama kondisi bernilai true, blok kode terus dijalankan.  
   Contoh: i <= 5

3. Post statement  
   Mengubah nilai variabel setelah setiap satu kali perulangan selesai.  
   Contoh: i++

### 2. Perbedaan For dengan 3 komponen dan For hanya menggunakan kondisi

For dengan tiga komponen memiliki inisialisasi, kondisi, dan post statement.

```go
for i := 1; i <= 5; i++ {
 fmt.Println(i)
}
```

For dengan satu kondisi bekerja seperti while, yaitu mengulang selama kondisi bernilai benar.

```go
i := 1
for i <= 5 {
 fmt.Println(i)
 i++
}
```

### 3. ganjil.go

```go
package main

import "fmt"

func main() {
	var n, jumlah int

	fmt.Scan(&n)

	for i := 1; i <= n; i += 2 {
		fmt.Print(i, " ")
		jumlah += i
	}

	fmt.Println()
	fmt.Println(jumlah)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/ganjil/output.png)

#### Deskripsi
Program ini berfungsi untuk mengetahui bilangan ganjil dari bilangan yang di inputkan atau dari 1 sampai n menggunakan for.
### 4. perkalian.go

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 1; i <= b; i++ {
		fmt.Printf("%d x %d = %d\n", a, i, a*i)
	}
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/perkalian/output.png)

#### Deskripsi
Program ini berfungsi sebagai kalkulator perkalian menggunakan for dan hasil perkalian di tulis semua dalam output.

## Kesimpulan
Kedua program tersebut menggunakan for untuk eksekusi kode yang ada di dalamnya selama kondisi terpenuhi.