package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fakX := faktorial(x)
	fakY := faktorial(y)

	n, r := x, y
	if x < y {
		n = y
		r = x
	}

	kom := faktorial(n) / (faktorial(r) * faktorial(n-r))
	fmt.Println(fakX, fakY, kom)
}
