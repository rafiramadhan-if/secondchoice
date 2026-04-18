package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Print("Masukkan elemen array: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// a. 
	fmt.Print("a. Semua elemen: ")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// b. 
	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// c. 
	fmt.Print("c. Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// d. 
	var x int
	fmt.Print("d. Masukkan nilai x (kelipatan): ")
	fmt.Scan(&x)
	fmt.Print("   Indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// e. 
	var hapus int
	fmt.Print("e. Masukkan indeks yang dihapus: ")
	fmt.Scan(&hapus)
	fmt.Print("   Array setelah hapus indeks ", hapus, ": ")
	for i := 0; i < n; i++ {
		if i != hapus {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// f. 
	sum := 0
	for _, v := range arr {
		sum += v
	}
	rataRata := float64(sum) / float64(n)
	fmt.Printf("f. Rata-rata: %.2f\n", rataRata)

	// g. 
	variansi := 0.0
	for _, v := range arr {
		variansi += math.Pow(float64(v)-rataRata, 2)
	}
	stdDev := math.Sqrt(variansi / float64(n))
	fmt.Printf("g. Standar deviasi: %.2f\n", stdDev)

	// h.
	var cari int
	fmt.Print("h. Masukkan bilangan yang dicari frekuensinya: ")
	fmt.Scan(&cari)
	frekuensi := 0
	for _, v := range arr {
		if v == cari {
			frekuensi++
		}
	}
	fmt.Printf("   Frekuensi %d: %d\n", cari, frekuensi)
}
