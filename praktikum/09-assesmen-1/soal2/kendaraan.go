package main

import "fmt"

func main() {
	var jenis string
	var jumlahM, jumlahK, jumlahT int

	for {
		_, gagal := fmt.Scan(&jenis)
		if gagal != nil {
			break
		}
		if jenis == "M" {
			jumlahM++
		} else if jenis == "K" {
			jumlahK++
		} else if jenis == "T" {
			jumlahT++
		} else {
			break
		}
	}

	fmt.Printf("Tipe M = %d\n", jumlahM)
	fmt.Printf("Tipe K = %d\n", jumlahK)
	fmt.Printf("Tipe T = %d\n", jumlahT)
}
