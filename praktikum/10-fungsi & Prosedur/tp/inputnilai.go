package main

import "fmt"

const NMAX int = 100

type Mahasiswa struct {
	nim, nama string
	nilai     float64
}
type TabMhs [NMAX]Mahasiswa

func cariNIM(T TabMhs, n int, targetNIM string) int {
	for i := 0; i < n; i++ {
		if T[i].nim == targetNIM {
			return i
		}
	}
	return -1
}

func rataRata(T TabMhs, n int) float64 {
	if n == 0 {
		return 0
	}
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += T[i].nilai
	}
	return total / float64(n)
}

func main() {
	var dataMhs TabMhs
	var jumlah int = 0
	var id, nama string
	var skor float64

	for {
		fmt.Scan(&id)
		if id == "STOP" {
			break
		}
		fmt.Scan(&nama, &skor)

		if cariNIM(dataMhs, jumlah, id) == -1 && jumlah < NMAX {
			dataMhs[jumlah].nim = id
			dataMhs[jumlah].nama = nama
			dataMhs[jumlah].nilai = skor
			jumlah++
		}
	}

	for i := 0; i < jumlah; i++ {
		fmt.Println(dataMhs[i].nim, dataMhs[i].nama, dataMhs[i].nilai)
	}
	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(dataMhs, jumlah))
}
