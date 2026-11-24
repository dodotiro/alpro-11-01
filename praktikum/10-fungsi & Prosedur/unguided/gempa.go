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
