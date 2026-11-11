package main

import "fmt"

func main() {
	var sebelumnya, sekarang int
	var tertinggi, terendah, jumlah, banyakData int
	var rataRata float64

	fmt.Scan(&sebelumnya)

	tertinggi = sebelumnya
	terendah = sebelumnya
	jumlah = sebelumnya
	banyakData = 1

	for {
		fmt.Scan(&sekarang)
		if sekarang == sebelumnya {
			break
		}
		if sekarang > tertinggi {
			tertinggi = sekarang
		}
		if sekarang < terendah {
			terendah = sekarang
		}
		jumlah += sekarang
		banyakData++
		sebelumnya = sekarang
	}

	rataRata = float64(jumlah) / float64(banyakData)
	fmt.Printf("%d %d %.2f\n", tertinggi, terendah, rataRata)
}
