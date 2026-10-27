package main

import "fmt"

func main()  {
	var bulanSebelumnya, bulanIni float64

	fmt.Scan(&bulanSebelumnya, &bulanIni)

	if bulanIni > bulanSebelumnya {
		selisih := bulanIni - bulanSebelumnya
		fmt.Printf("Peningkatan ssebesar %g\n", selisih)
	} else if bulanIni < bulanSebelumnya {
		selisih := bulanSebelumnya - bulanIni
		fmt.Printf("Penurunan sebesar %g\n", selisih)
	} else {
		fmt.Println("Tetap")
	}
}