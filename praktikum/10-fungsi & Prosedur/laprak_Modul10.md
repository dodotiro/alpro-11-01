# <h1 align="center">Laporan Praktikum Modul 10 - Fungsi & Prosedur</h1>
<p align="center">Nur Widodo - 109092630005</p>

## Dasar Teori

### A. Fungsi
Fungsi adalah blok kode terisolasi yang dapat digunakan kembali untuk menjalankan tugas tertentu, sehingga program lebih terstruktur, mudah dibaca, diuji, dan dipelihara. Suatu subprogram umumnya ditandai sebagai fungsi apabila:

Memiliki deklarasi tipe nilai kembali (return type).

Memuat pernyataan return di dalam badan fungsi untuk mengembalikan hasil proses ke pemanggil.

### B. Parameter dan Pointer
- Parameter: Variabel lokal di dalam tanda kurung fungsi (func(param Tipe)) untuk menampung input saat fungsi dipanggil. Secara default, Go menerapkan pass by value, yaitu menyalin nilai ke alamat memori baru sehingga perubahan di dalam fungsi tidak mengubah variabel asli di luar fungsi (kecuali lewat pointer).

- Pointer: Variabel yang menyimpan alamat memori dari suatu nilai, bukan nilainya langsung. Memanfaatkan pointer membantu menghemat memori saat mengirim data besar dan memungkinkan fungsi mengubah nilai variabel aslinya. Operator utamanya meliputi:
	1. Address-of (&): Mengambil alamat memori dari sebuah variabel.

	2. Dereference (*): Mendeklarasikan tipe pointer (*T) atau membaca/mengubah nilai pada alamat memori yang ditunjuk.

### C. Array dan Struct
- Array: Kumpulan elemen bertipe data sama dengan panjang tetap (fixed size) yang ditentukan saat deklarasi dan tidak dapat diubah ukurannya secara dinamis.

- Struct: Tipe data bentukan yang menggabungkan beberapa field bertipe data berbeda ke dalam satu kesatuan untuk merepresentasikan objek atau entitas yang lebih kompleks.

## Guided

### 1. sapi.go

```go
package main

import "fmt"

const NMAX = 10000

type Arr [NMAX]int

func mengisiSapi(T *Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
}

func geserKanan(T *Arr, n int) {
	var temp int = T[n-1]

	for i := n - 1; i > 0; i-- {
		T[i] = T[i-1]
	}

	T[0] = temp
}

func cetak(T Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var T Arr
	mengisiSapi(&T, n)

	k = k % n

	for i := 0; i < k; i++ {
		geserKanan(&T, n)
	}
	cetak(T, n)
}

```

#### Deskripsi
Program ini untuk memindahkan sapi dengan aturan sapi di kandang terakhir pindah ke kandang pertama, sedangkan sapi lainnya bergeser satu kandang ke kanan.

### 2. player.go

```go
package main

import "fmt"

const NMAX int = 999

type player struct {
	nickname string
	score    int
}

type tabPlayer [NMAX]player

func rataRata(T tabPlayer, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].score
	}
	return float64(total) / float64(n)
}

func idxTertinggi(T tabPlayer, n int) int {
	idx := 0
	for i := 1; i < n; i++ {
		if T[i].score > T[idx].score {
			idx = i
		}
	}
	return idx
}

func hitungDiAtas(T tabPlayer, n int, batas float64) int {
	count := 0
	for i := 0; i < n; i++ {
		if float64(T[i].score) > batas {
			count++
		}
	}
	return count
}

func main() {
	var T tabPlayer
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nickname, &T[i].score)
	}

	avg := rataRata(T, n)
	best := idxTertinggi(T, n)
	di := hitungDiAtas(T, n, avg)

	fmt.Printf("Rata-rata skor: %.2f\n", avg)
	fmt.Printf("Player terbaik: %s (%d)\n", T[best].nickname, T[best].score)
	fmt.Printf("Jumlah player di atas rata-rata: %d\n", di)
}

```

#### Deskripsi
Program ini untuk mengolah skor pemain sebuah turnamen game dan menghitung rata-rata skor, mencari pemain dengan skor tertinggi, dan mencari jumlah pemain yang skornya diatas rata-rata.

## Unguided

### 1. gempa.go

```go
package main

import "fmt"

const NMAX int = 999

type gempa struct {
	nama      string
	tahun     int
	magnitudo float64
}

type tabGempa [NMAX]gempa

func kategori(mag float64) string {
	if mag < 4.0 {
		return "Minor"
	} else if mag < 6.0 {
		return "Sedang"
	} else if mag < 7.0 {
		return "Kuat"
	}
	return "Besar"
}

func rataMagnitudo(T tabGempa, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += T[i].magnitudo
	}
	return total / float64(n)
}

func idxTerkuat(T tabGempa, n int) int {
	maxIdx := 0
	for i := 1; i < n; i++ {
		if T[i].magnitudo > T[maxIdx].magnitudo {
			maxIdx = i
		}
	}
	return maxIdx
}

func hitungKategori(T tabGempa, n int, kat string) int {
	var total int
	for i := 0; i < n; i++ {
		if kategori(T[i].magnitudo) == kat {
			total++
		}
	}
	return total
}

func main() {
	var n, terkuat, jmlBesar int
	var data tabGempa

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i].nama, &data[i].tahun, &data[i].magnitudo)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %.1f %s\n", data[i].nama, data[i].tahun, data[i].magnitudo, kategori(data[i].magnitudo))
	}

	rata := rataMagnitudo(data, n)
	terkuat = idxTerkuat(data, n)
	jmlBesar = hitungKategori(data, n, "Besar")

	fmt.Printf("Rata-rata magnitudo: %.2f\n", rata)
	fmt.Printf("Gempa terkuat: %s %d %.1f\n", data[terkuat].nama, data[terkuat].tahun, data[terkuat].magnitudo)
	fmt.Println("Jumlah gempa kategori Besar:", jmlBesar)
}

```

##### Output
![Screenshot Output Unguided](./Unguided/gempa.png)

#### Deskripsi
Program ini mengolah data gempa melalui fungsi kategori (klasifikasi magnitudo), rataMagnitudo (rata-rata magnitudo), idxTerkuat (indeks gempa terbesar), dan hitungKategori (jumlah gempa per kategori dari "Minor" hingga "Besar"), yang seluruh alur input-output dan pemrosesannya dikendalikan oleh main().


### 2. nim.go

```go
package main

import "fmt"

const NMAX int = 2002

type tabNIM [NMAX]int

func cariNIM(T tabNIM, n, target int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i] == target {
			*idx = i
			break
		}
	}
}

func ubahNIM(T *tabNIM, idx, nimBaru int) {
	T[idx] = nimBaru
}

func cetak(T tabNIM, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i])
	}
}

func main() {
	var data tabNIM
	var n, target, baru, posisi int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
	fmt.Scan(&target)

	cariNIM(data, n, target, &posisi)
	if posisi == -1 {
		fmt.Println("NIM tidak ditemukan dalam sistem")
	} else {
		fmt.Scan(&baru)
		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", target, posisi+1)
		ubahNIM(&data, posisi, baru)
		fmt.Println("Daftar NIM setelah terjadi perubahan data:")
		cetak(data, n)
	}
}

```

##### Output
![Screenshot Output Unguided](./Unguided/nim.png)

#### Deskripsi
Program ini mencari NIM mahasiswa di dalam array lalu menggantinya jika ditemukan melalui cariNIM (mencari NIM x pada n elemen array T dan menyimpan indeks pertamanya ke idx, atau -1 jika tidak ada), ubahNIM (mengubah nilai pada indeks idx di array pointer T menjadi baru), dan cetak (menampilkan elemen array T dari indeks 0 hingga n-1).

## Kesimpulan
1. Fungsi membagi kode menjadi bagian modular yang dapat dipakai ulang, serta umumnya memiliki deklarasi nilai kembali dan return.

2. Parameter menerima data masuk secara pass by value, sedangkan pointer menyimpan alamat memori dengan operator & (address-of) dan * (dereference).

3. Array menyimpan data bertipe sama dengan ukuran tetap, sedangkan struct mengelompokkan data bertipe berbeda dalam satu struktur.

## Referensi
1. Go Team. (2026). _The Go Programming Language Specification_. Google LLC. Diakses pada 04 Oktober 2026 melalui https://go.dev/ref/spec.
