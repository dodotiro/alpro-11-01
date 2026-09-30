package main

import "fmt"

func main() {
	var c float64

	fmt.Println("Masukkan suhu celcius")
	fmt.Scan(&c)

	r := (4.0 / 5.0) * c
	fmt.Println(r, "Reamur")

}
