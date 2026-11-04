package main

import "fmt"

func main() {
	var tanggal, bulan, tahun int

	fmt.Scan(&tanggal, &bulan, &tahun)

	kabisat := tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0)

	jumlahHari := 0

	if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 ||
		bulan == 8 || bulan == 10 || bulan == 12 {
		jumlahHari = 31
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		jumlahHari = 30
	} else if bulan == 2 {
		if kabisat {
			jumlahHari = 29
		} else {
			jumlahHari = 28
		}
	}

	if bulan < 1 || bulan > 12 || tanggal < 1 || tanggal > jumlahHari {
		fmt.Println("Tanggal tidak valid")
		return
	}

	fmt.Println("Valid")

	tanggal++

	if tanggal > jumlahHari {
		tanggal = 1
		bulan++

		if bulan > 12 {
			bulan = 1
			tahun++
		}
	}

	fmt.Printf("%d/%d/%d\n", tanggal, bulan, tahun)
}
