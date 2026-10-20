package main

import "fmt"

func main() {
	lamp_status := "kuning"

		if lamp_status == "merah" {
			fmt.Println("Berhenti")
		} else if lamp_status == "kuning" {
			fmt.Println ("Hati-hati / Lebih pelan")
		} else if lamp_status == "hijau" {
			fmt.Println ("jalan")
		} else {
			fmt.Print("Perintah tidak terdaftar disini")
		}
}