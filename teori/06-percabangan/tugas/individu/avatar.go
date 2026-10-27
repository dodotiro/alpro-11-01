package main

import (
	"fmt"
)

func main() {
	var orang int
	fmt.Scan(&orang)

	dewasa := orang / 5
	if dewasa > 3 {
		dewasa = 3
	}
	orang -= dewasa * 5

	kecil := orang / 2
	if kecil > 5 {
		kecil = 5
	}
	orang -= kecil * 2

	takBerangkat := orang

	if dewasa > 0 {
		fmt.Printf("Dewasa %d", dewasa)
	}
	if kecil > 0 {
		if dewasa > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("kecil %d", kecil)
	}
	if takBerangkat > 0 {
		if dewasa > 0 || kecil > 0 {
			fmt.Print(", dan ")
		}
		fmt.Printf("%d tak berangkat", takBerangkat)
	}
	fmt.Println()
}