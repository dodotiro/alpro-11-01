package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x == 0 {
		fmt.Println("Nol")
	} else if x%2 == 0 {
		fmt.Println("Genap")
		if x < 0 {
			fmt.Println("Negatif")
		} else if x >= 0 {
			fmt.Println("Positif")
		}
	} else if x%2 != 0 {
		fmt.Println("Ganjil")
		if x < 0 {
			fmt.Print("Negatif")
		} else if x >= 0 {
			fmt.Println("Positif")
		}
	}
}
