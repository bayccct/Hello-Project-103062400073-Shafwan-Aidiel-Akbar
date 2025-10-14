package main

import (
	"fmt"
)

// Tugas Smester 2
func main() {
	var kendaraan string
	var waktu int
	fmt.Print("Masukan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)

	fmt.Print("Masukan durasi parkir (dalam menit): ")
	fmt.Scan(&waktu)
	switch kendaraan {
	case "Motor":
		tarifPerJam1 := 1500
		if waktu >= 30 {
			sisa := waktu % 60
			tes := waktu - sisa
			sisa *= 100
			tarifPerJam1 *= tes / 60
			fmt.Println(sisa + tarifPerJam1)
		} else if waktu < 30 {
			sisa := waktu % 60
			tes := waktu - sisa
			sisa *= 200
			tarifPerJam1 *= tes / 60
			fmt.Println(sisa + tarifPerJam1)
		} else {
			sisa := waktu % 60
			waktu -= sisa
			fmt.Println("Sisa waktu parkir gratis karena lebih daari 5 jam.")
			fmt.Println(tarifPerJam1 * (waktu / 60))
		}

	case "Mobil":
		tarifPerJam2 := 2000
		if waktu >= 30 {
			sisa := waktu % 60
			tes := waktu - sisa
			sisa *= 100
			tarifPerJam2 *= tes / 60
			fmt.Println(sisa + tarifPerJam2)
			return
		} else if waktu < 30 {
			sisa := waktu % 60
			tes := waktu - sisa
			sisa *= 200
			tarifPerJam2 *= tes / 60
			fmt.Println(sisa + tarifPerJam2)
			return
		} else {
			sisa := waktu % 60
			waktu -= sisa
			fmt.Println("Sisa waktu parkir gratis karena lebih daari 5 jam.")
			fmt.Println(tarifPerJam2 * (waktu / 60))
			return
		}

	case "Truk":
		tarifPerJam3 := 3000
		if waktu >= 30 {
			sisa := waktu % 60
			tes := waktu - sisa
			sisa *= 100
			tarifPerJam3 *= tes / 60
			fmt.Println(sisa + tarifPerJam3)
			return
		} else if waktu < 30 {
			sisa := waktu % 60
			tes := waktu - sisa
			sisa *= 200
			tarifPerJam3 *= tes / 60
			fmt.Println(sisa + tarifPerJam3)
			return
		} else if waktu > 300 {
			sisa := waktu % 60
			tes := waktu - sisa
			sisa *= 0
			tarifPerJam3 *= tes / 60
			fmt.Println(sisa + tarifPerJam3)
			return
		}

	default:
		fmt.Println("hanya ada motor, mobil, dan truk")
	}

}
