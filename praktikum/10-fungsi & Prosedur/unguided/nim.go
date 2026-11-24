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
