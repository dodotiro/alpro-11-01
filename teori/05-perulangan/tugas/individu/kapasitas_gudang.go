package main

import "fmt"

func main() {
	var kapasitas, pengiriman int
	fmt.Scan(&kapasitas)

	totalBarang := 0

	for i := 0; i < 5; i++ {
		fmt.Scan(&pengiriman)
		totalBarang += pengiriman

		if totalBarang >= kapasitas {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	}
}
