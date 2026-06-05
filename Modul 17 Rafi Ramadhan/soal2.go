package main

import "fmt"

func main() {
	var x, s string
	var n int

	fmt.Print("String yang dicari: ")
	fmt.Scan(&x)

	fmt.Print("Jumlah string: ")
	fmt.Scan(&n)

	jumlah := 0
	posisi := -1

	for i := 1; i <= n; i++ {
		fmt.Scan(&s)

		if s == x {
			jumlah++

			if posisi == -1 {
				posisi = i
			}
		}
	}

	fmt.Println("Ada?", jumlah > 0)

	if posisi != -1 {
		fmt.Println("Posisi pertama:", posisi)
	} else {
		fmt.Println("Tidak ditemukan")
	}

	fmt.Println("Jumlah kemunculan:", jumlah)
	fmt.Println("Minimal dua kali?", jumlah >= 2)
}
