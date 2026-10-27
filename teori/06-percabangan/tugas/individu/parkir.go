package main

import (
	"fmt"
)

func main() {
	var h1, m1, h2, m2 int
	fmt.Scan(&h1, &m1, &h2, &m2)

	if h1 >= 1 && h1 <= 5 {
		h1 += 12
	}
	if h2 >= 1 && h2 <= 5 {
		h2 += 12
	}

	mulai := (h1-7)*60 + m1
	selesai := (h2-7)*60 + m2

	durasi := selesai - mulai
	hh := durasi / 60
	mm := durasi % 60

	fmt.Printf("%d jam %d menit\n", hh, mm)
}