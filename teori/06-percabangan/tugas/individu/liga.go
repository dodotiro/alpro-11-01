package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	maks := a
	if b > maks {
		maks = b
	}
	if c > maks {
		maks = c
	}
	if d > maks {
		maks = d
	}

	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	if d < min {
		min = d
	}

	fmt.Printf("%d %d\n", maks, min)
}