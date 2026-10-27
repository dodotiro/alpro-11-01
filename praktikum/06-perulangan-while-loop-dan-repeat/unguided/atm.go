package main

import "fmt"

func main() {

	var passwordBenar string = "1234"
	var saldo int = 10000000
	var passwordInput string
	var kesempatan int = 3
	var loginBerhasil bool = false

	for kesempatan > 0 {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&passwordInput)

		if passwordInput == passwordBenar {
			loginBerhasil = true
			fmt.Println("Login berhasil!")
			break
		} else {
			kesempatan--
			if kesempatan > 0 {
				fmt.Println("Password salah. Silakan coba lagi.")
			}
		}
	}

	if !loginBerhasil {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	var pilihan int = 0
	for pilihan != 3 {
		fmt.Println("\nMenu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		case 2:
			var jumlahTarik int
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
			fmt.Scan(&jumlahTarik)

			if jumlahTarik <= saldo {
				saldo -= jumlahTarik
				fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\n", jumlahTarik)
				fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		case 3:
			fmt.Println("Keluar dari sistem. Terima kasih!")
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu 1, 2, atau 3.")
		}
	}
}
