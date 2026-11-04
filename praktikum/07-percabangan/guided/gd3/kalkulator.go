package main

import "fmt"

func main() {
	var a, b int
	var op string

	fmt.Scan(&a, &op, &b)

	switch op {
	case "+":
		fmt.Print(a + b)
	case "-":
		fmt.Print(a - b)
	case "*":
		fmt.Print(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Error: Pembagian dengan nol")
		} else {
			fmt.Printf("%.2f\n", float64(a)/float64(b))
		}
	case "%":
		if b == 0 {
			fmt.Println("Error: Pembagian dengan nol")
		} else {
			fmt.Println(a % b)
		}
	default:
		fmt.Println("Operator tidak valid")
	}
}
