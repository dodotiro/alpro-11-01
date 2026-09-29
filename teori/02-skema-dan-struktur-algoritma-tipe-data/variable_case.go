package main

import "fmt"

func main() {
	var name string

	name = "Yudha Islami Sulistya"
	fmt.Println("Nama : ", name)

	// name = 17
	// fmt.Println(name)
	//akan error karena string di kasih nilai int = 17

	var lastName = "Sulistya"
	fmt.Println("Nama Belakang : ", lastName)

	middleName := "Islami"
	fmt.Println("Nama Tengah: ", middleName)

	var (
		fullName  = "Yudha Islami Sulistya"
		firstName = "Yudha"
	)

	fmt.Println(fullName)
	fmt.Println(firstName)

}
