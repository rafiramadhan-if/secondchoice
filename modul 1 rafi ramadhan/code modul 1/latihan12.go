package main

import "fmt"

func main() {
	var b, jumlahFaktor int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}
	fmt.Println()

	fmt.Printf("Prima: %t\n", jumlahFaktor == 2)
}
