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
