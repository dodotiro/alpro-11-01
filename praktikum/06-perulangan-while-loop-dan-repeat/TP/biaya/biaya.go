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
