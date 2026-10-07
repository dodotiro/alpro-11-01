package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := n / 10
	b := n % 10

	fmt.Println((11 * a * 100) + (11 * b))
}
