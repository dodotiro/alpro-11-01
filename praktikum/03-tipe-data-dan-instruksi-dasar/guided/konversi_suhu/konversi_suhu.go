package main

import "fmt"

func main() {
	var c float64

	fmt.Println("Masukkan suhu celcius")
	fmt.Scan(&c)

	k := c + 273
	fmt.Println(k, "kelvin")

}
