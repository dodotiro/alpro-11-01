package main

import "fmt"

const NMAX int = 999

type Raport struct {
	matkul string
	skor   int
}
type TabRaport [NMAX]Raport

func konversiHuruf(n int) string {
	if n > 80 && n <= 100 {
		return "A"
	} else if n > 70 && n <= 80 {
		return "AB"
	} else if n > 65 && n <= 70 {
		return "B"
	} else if n > 60 && n <= 65 {
		return "BC"
	} else if n > 50 && n <= 60 {
		return "C"
	} else if n > 40 && n <= 50 {
		return "D"
	}
	return "E"
}

func rataRata(T TabRaport, n int) float64 {
	var total int = 0
	for i := 1; i <= n; i++ {
		total += T[i].skor
	}
	return float64(total) / float64(n)
}

func main() {
	var listNilai TabRaport
	var jumlah int
	fmt.Scan(&jumlah)

	for i := 1; i <= jumlah; i++ {
		fmt.Scan(&listNilai[i].matkul, &listNilai[i].skor)
	}
	for i := 1; i <= jumlah; i++ {
		fmt.Printf("%s %d %s\n", listNilai[i].matkul, listNilai[i].skor, konversiHuruf(listNilai[i].skor))
	}
	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(listNilai, jumlah))
}
