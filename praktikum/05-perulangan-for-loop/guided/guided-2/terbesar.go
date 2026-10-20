package main

import "fmt"

func main() {
	var n, angka, max int
	fmt.Scan(&n, &angka)

	max = angka
	for i := 2; i <= n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}
	}
	fmt.Println(max)
}
