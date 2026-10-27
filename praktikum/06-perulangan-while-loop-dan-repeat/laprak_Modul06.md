# <h1 align="center">Laporan Praktikum Modul 06 - Perulangan While</h1>
<p align="center">Nur Widodo - 109092630005</p>

## Dasar Teori

### A. Pengenalan while
 Bahasa Go tidak memiliki kata kunci while maupun do-while. Perulangan while ditulis menggunakan for yang hanya berisi kondisi: 

 ```
 for kondisi { 
	// perintah yang diulang selama kondisi bernilai true 
	} 
 ```

## Guided

### 1. terkecil.go

```go
package main

import "fmt"

func main() {
	var batas int
	fmt.Scan(&batas)

	n := 1
	for n*n*n <= batas {
		n++
	}

	fmt.Println(n)
}

```

#### Deskripsi
Program ini mencari bilangan bulat terkecil n sehingga n^3 lebih besar dari nilai yang di inputkan.

### 2. terbesar.go

```go
package main

import "fmt"

func main() {
	var batas int
	fmt.Scan(&batas)

	n := 1
	for n*n < batas {
		n++
	}

	fmt.Println(n - 1)
}

```

#### Deskripsi
Program ini mencari bilangan bulat terbesar n sehingga n^2 kurang dari nilai yang di inputkan.

### 3. vokal.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	kalimat, _ := reader.ReadString('\n')

	i := 0
	jumlahVokal := 0

	for i < len(kalimat) {
		ch := strings.ToLower(string(kalimat[i]))
		if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
			jumlahVokal++
		}
		i++
	}

	fmt.Println(jumlahVokal)
}

```

#### Deskripsi
Program ini berfungsi menghitung berapa banyak huruf vokal (a, i, u, e ,o) yang terdapat dalam kata/kalimat.

## Unguided

### 1. atm.go

```go
package main

import "fmt"

func main() {

	var passwordBenar string = "1234"
	var saldo int = 10000000
	var passwordInput string
	var kesempatan int = 3
	var loginBerhasil bool = false

	for kesempatan > 0 {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&passwordInput)

		if passwordInput == passwordBenar {
			loginBerhasil = true
			fmt.Println("Login berhasil!")
			break
		} else {
			kesempatan--
			if kesempatan > 0 {
				fmt.Println("Password salah. Silakan coba lagi.")
			}
		}
	}

	if !loginBerhasil {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	var pilihan int = 0
	for pilihan != 3 {
		fmt.Println("\nMenu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		case 2:
			var jumlahTarik int
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
			fmt.Scan(&jumlahTarik)

			if jumlahTarik <= saldo {
				saldo -= jumlahTarik
				fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\n", jumlahTarik)
				fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		case 3:
			fmt.Println("Keluar dari sistem. Terima kasih!")
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu 1, 2, atau 3.")
		}
	}
}

```

##### Output
![Screenshot Output Unguided](./Unguided/output.png)

#### Deskripsi
Program ini layaknya atm yang menampilkan masukkan pin, lihat saldo dan tarik tunai.

## Kesimpulan
Semua program diatas menggunakan fungsi perulangan while:
 ```
 for kondisi { 
	// perintah yang diulang selama kondisi bernilai true 
	} 
 ```


## Referensi
1. Go Team. (2026). _The Go Programming Language Specification_. Google LLC. Diakses pada 04 Oktober 2026 melalui https://go.dev/ref/spec.
