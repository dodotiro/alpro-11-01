package main

import "fmt"

func main() {
	var totalHari int
	fmt.Print("Masukkan Jumlah Hari: ")
	fmt.Scan(&totalHari)

	hariPertahun := 360
	hariPerbulan := 30
	hariPerminggu := 7

	tahun := totalHari / hariPertahun
	sisaHari := totalHari % hariPertahun

	bulan := sisaHari / hariPerbulan
	sisaHari = sisaHari % hariPerbulan

	minggu := sisaHari / hariPerminggu
	hari := sisaHari % hariPerminggu

	fmt.Printf("%d hari setara dengan: %d tahun, %d bulan, %d minggu, dan %d hari.\n", totalHari, tahun, bulan, minggu, hari)
}
