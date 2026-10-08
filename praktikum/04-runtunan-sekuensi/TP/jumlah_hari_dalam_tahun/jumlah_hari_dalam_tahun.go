package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Println("Masukkan tahun: ")
	fmt.Scan(&tahun)
	fmt.Println("Masukkan bulan: ")
	fmt.Scan(&bulan)

	bulanKabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println("Jumlah hari: 31")
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println("Jumlah hari: 30")
	case "Feb":
		if bulanKabisat {
			fmt.Println("Jumlah hari: 29")
		} else {
			fmt.Println("Jumlah hari: 28")
		}
	default:
		fmt.Println("-")
	}
}
