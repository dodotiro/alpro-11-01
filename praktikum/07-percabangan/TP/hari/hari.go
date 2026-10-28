package main

import "fmt"

func main() {
	var hari int
	fmt.Scan(&hari)

	switch {
	case hari == 1:
		fmt.Println("Senin")
	case hari == 2:
		fmt.Println("Selasa")
	case hari == 3:
		fmt.Println("Rabu")
	case hari == 4:
		fmt.Println("Kamis")
	case hari == 5:
		fmt.Println("Jumat")
	case hari == 6:
		fmt.Println("Sabtu", "Akhir Pekan")
	case hari == 7:
		fmt.Println("Minggu", "Akhir Pekan")
	default:
		fmt.Println("Nomor hari tidak valid")
	}
}
