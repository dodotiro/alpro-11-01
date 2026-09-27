package main

import "fmt"

func main() {
	var a, b int
	var operator string

	//Input angka pertama
	fmt.Print("Masukkan Angka Pertama :")
	fmt.Scan(&a)

	//Pilih operator
	fmt.Print("Masukkan Operator (+, -, *, /, %): ")
	fmt.Scan(&operator)

	//Input angka kedua
	fmt.Print("Masukkan Angka Kedua :")
	fmt.Scan(&b)

	//Menampilkan hasil
	switch operator {
	case "+":
		fmt.Printf("Hasil: %d + %d = %d\n", a, b, a+b)
	case "-":
		fmt.Printf("Hasil: %d - %d = %d\n", a, b, a-b)
	case "*":
		fmt.Printf("Hasil: %d * %d = %d\n", a, b, a*b)
	case "/":
		if b == 0 {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan!")
		} else {
			fmt.Printf("Hasil: %d / %d = %d\n", a, b, a/b)
		}
	case "%":
		if b == 0 {
			fmt.Println("Error: Modulus dengan nol tidak diperbolehkan!")
		} else {
			fmt.Printf("Hasil: %d %% %d = %d\n", a, b, a%b) // Gunakan %% untuk mencetak simbol % di Printf
		}
	default:
		fmt.Println("Operator tidak dikenal!")
	}
}
