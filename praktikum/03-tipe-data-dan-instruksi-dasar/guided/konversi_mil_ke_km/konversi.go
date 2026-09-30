package main

import "fmt"

func main() {
	var mil float64

	fmt.Println("Masukkan jarak mil")
	fmt.Scan(&mil)

	fmilkm := mil * 1.6
	fmt.Println(fmilkm, "km")

}
