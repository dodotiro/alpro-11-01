package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a, b := n, m
	for b != 0 {
		a, b = b, a%b
	}

	kpk := (n * m) / a
	fmt.Println(kpk)
}
