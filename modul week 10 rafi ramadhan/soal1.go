package main

import "fmt"

func main() {
	const Nmax = 1000
	var berat [Nmax]float64
	var n int

	fmt.Print("Masukkan jumlah data berat: ")
	fmt.Scan(&n)

	if n <= 0 || n > Nmax {
		fmt.Println("Jumlah data tidak valid!")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Println()
	fmt.Printf("Berat minimum: %.2f\n", min)
	fmt.Printf("Berat maksimum: %.2f\n", max)
}
