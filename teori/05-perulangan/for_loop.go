package main

import "fmt"

func main() {
	var n, i int
	fmt.Scan(&n)
	i = n
	for n = 1; n <= i; n++ {
		fmt.Println(n)
	}
}
