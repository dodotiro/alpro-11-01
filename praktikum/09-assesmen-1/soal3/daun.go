package main

import "fmt"

func main() {
	var jumlahDaun, lebar, terkecil int
	fmt.Scan(&jumlahDaun)

	for i := 0; i < jumlahDaun; i++ {
		fmt.Scan(&lebar)
		if i == 0 || lebar < terkecil {
			terkecil = lebar
		}
	}

	fmt.Println(terkecil)
}
