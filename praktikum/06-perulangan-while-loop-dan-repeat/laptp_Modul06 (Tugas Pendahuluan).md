# <h1 align="center">Tugas Pendahuluan Modul 06 - Perulangan While</h1>
<p align="center">Nur Widodo - 109092630005</p>

### 1. lulus.go

```go
package main

import "fmt"

func main() {
	var nilai int

	for {
		fmt.Print("Masukan nilai Anda: ")
		fmt.Scan(&nilai)

		if nilai == -1 {
			fmt.Println("Program selesai.")
			break
		}

		if nilai >= 60 {
			fmt.Println("Anda lulus ujian.")
		} else {
			fmt.Println("Anda tidak lulus ujian.")
		}
	}
}

```

#### Output
![Screenshot Output Unguided](./tp/lulus/output.png)


#### Deskripsi
Program ini menentukan siswa yang lulus dan tidak lulus, dimana nilai >= 60 adalah lulus dan jika ada nilai < 0 maka program selesai.

### 2. deret.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	jumlah := 0
	i := 1

	for i <= n {
		jumlah += i
		i++
	}

	fmt.Println(jumlah)
}

```

#### Output
![Screenshot Output Unguided](./tp/deret/output.png)


#### Deskripsi
Program untuk menghitung jumlah dari sederet angka yang dimasukkan menggunakan perulangan dimana pengitungan dimulai dari 1 hingga n.
### 3. duaterendah.go

```go
package main

import "fmt"

func main() {
	var jumlahSiswa int
	fmt.Scan(&jumlahSiswa)

	var namaTerendah, namaKeduaTerendah string
	nilaiTerendah := 999999
	nilaiKeduaTerendah := 999999

	i := 0
	for i < jumlahSiswa {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if nilai < nilaiTerendah {
			nilaiKeduaTerendah = nilaiTerendah
			namaKeduaTerendah = namaTerendah
			nilaiTerendah = nilai
			namaTerendah = nama
		} else if nilai < nilaiKeduaTerendah {
			nilaiKeduaTerendah = nilai
			namaKeduaTerendah = nama
		}

		i++
	}

	fmt.Println(namaTerendah)
	fmt.Println(namaKeduaTerendah)
}

```

##### Output
![Screenshot Output Unguided](./tp/duaterendah/output.png)


#### Deskripsi
Program ini berfungsi untuk menentukan 2 nilai terendah dari n siswa.

### 4. biaya.go

```go
package main

import "fmt"

func main() {
	var tahunKe int
	fmt.Scan(&tahunKe)

	var biayaKuliah float64 = 10000000

	tahun := 2
	for tahun <= tahunKe {
		biayaKuliah = biayaKuliah * 1.06
		tahun++
	}

	fmt.Printf("%.0f\n", biayaKuliah)
}

```
#### Output
![Screenshot Output Unguided](./tp/biaya/output.png)


#### Deskripsi
Program ini membantu sebuah universitas menghitung perkiraan biaya kuliah setiap tahun sampai tahun ke-n. Biaya kuliah awal pada universitas tersebut adalah Rp 10.000.000 dan setiap tahun bertambah sebesar 6% dari biaya tahun sebelumnya.

## Kesimpulan
Bahasa Go (Golang) tidak memiliki kata kunci while. Sebagai gantinya, fungsi while loop dibuat menggunakan kata kunci for yang hanya diisi dengan satu ekspresi kondisi saja (tanpa blok inisialisasi dan post-statement).