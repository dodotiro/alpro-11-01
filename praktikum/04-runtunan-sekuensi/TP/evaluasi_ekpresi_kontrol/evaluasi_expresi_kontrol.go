package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	if sngNum > 0 || (intNum >= 0 && -1*intOther == -10) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
