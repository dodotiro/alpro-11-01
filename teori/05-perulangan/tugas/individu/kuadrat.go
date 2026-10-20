package main

import (
	"fmt"
	"strings"
)

func main() {
	var x int
	fmt.Scan(&x)

	var kuadratList []string
	var total int

	for x > 0 {
		digit := x % 10
		k := digit * digit
		total += k
		kuadratList = append(kuadratList, fmt.Sprintf("%d", k))
		x /= 10
	}

	fmt.Println(strings.Join(kuadratList, " "))
	fmt.Println(total)
}
