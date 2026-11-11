package main

import "fmt"

func main() {
	var daun, jumlah int

	fmt.Scan(&jumlah)
	fmt.Scan(&daun)

	n := 0
	for n < daun {
		n++
		daun = n
	}

	fmt.Println(n)
}
