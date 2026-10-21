package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	jumlah := 0
	i := 1

	for i <= n {
		jumlah += i
		i++
	}

	fmt.Println(jumlah)
}
