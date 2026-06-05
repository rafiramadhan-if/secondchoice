package main

import "fmt"

func main() {
	var x, total float64
	var count int

	fmt.Println("Masukkan bilangan (9999 untuk berhenti):")
	fmt.Scan(&x)

	for x != 9999 {
		total += x
		count++
		fmt.Scan(&x)
	}

	if count == 0 {
		fmt.Println("Tidak ada data")
	} else {
		fmt.Printf("Rerata = %.2f\n", total/float64(count))
	}
}
