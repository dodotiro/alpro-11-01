package main

import "fmt"

func main() {
	var x, y int

	fmt.Println("Masukkan jumlah kue: ")
	fmt.Scan(&y)

	fmt.Println("Masukkan jumlah anggota keluarga: ")
	fmt.Scan(&x)

	sisa := y % x
	fmt.Println("Sisa kue adalah: ", sisa)

}
