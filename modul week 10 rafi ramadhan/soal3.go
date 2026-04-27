package main

import "fmt"

const Nmax = 100

type arrBalita [Nmax]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max float64

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	if n <= 0 || n > Nmax {
		fmt.Printf("Jumlah data harus antara 1 hingga %d!\n", Nmax)
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
		if data[i] <= 0 {
			fmt.Println("Berat balita harus lebih dari 0!")
			return
		}
	}

	hitungMinMax(data, n, &min, &max)
	rataRata := rerata(data, n)

	fmt.Println()
	fmt.Printf("Berat balita minimum : %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", max)
	fmt.Printf("Rerata berat balita  : %.2f kg\n", rataRata)
}
