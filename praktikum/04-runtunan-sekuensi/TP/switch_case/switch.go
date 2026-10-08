package main

import "fmt"

func main() {
	var os string

	fmt.Println("Masukkan Operasi Sistem: ")
	fmt.Scan(&os)

	switch os {
	case "windows", "macos":
		fmt.Println("Close Source")
	case "linux":
		fmt.Println("Open Source")
	default:
		fmt.Println("-")
	}
}
