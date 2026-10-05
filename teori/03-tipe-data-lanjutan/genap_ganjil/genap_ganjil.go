package main

import "fmt"

func main() {
	var n int
	var hasil boolean
	fmt.Scan(&n)

	hasil := n%2 == 0

	println(hasil)
}
