package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			count++
		}
	}
	fmt.Println()

	if count == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
