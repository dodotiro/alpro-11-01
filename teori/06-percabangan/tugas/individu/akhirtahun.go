package main

import (
	"fmt"
)

func main() {
	var totalBelanja float64
	var bersediaKartu bool

	fmt.Scan(&totalBelanja, &bersediaKartu)

	dapatKartu := bersediaKartu
	dapatDiskon := false
	dapatCashback := false

	if totalBelanja >= 100000 && dapatKartu {
		dapatDiskon = true
		totalBelanja = totalBelanja - (0.10 * totalBelanja)
	}

	if totalBelanja >= 200000 && dapatKartu {
		dapatCashback = true
		totalBelanja = totalBelanja - 75000
	}

	fmt.Printf("Kartu? %t\n", dapatKartu)
	fmt.Printf("Diskon? %t\n", dapatDiskon)
	fmt.Printf("Cashback? %t\n", dapatCashback)
	fmt.Printf("Rp. %.0f\n", totalBelanja)
}