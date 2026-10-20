package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	isCantik := true
	prevDigit := 10

	for x > 0 {
		currentDigit := x % 10
		if prevDigit < currentDigit {
			isCantik = false
			break
		}
		prevDigit = currentDigit
		x /= 10
	}

	fmt.Println(isCantik)
}
