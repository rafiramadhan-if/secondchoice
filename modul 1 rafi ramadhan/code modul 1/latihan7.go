package main

import "fmt"

func main() {
	var bunga string
	var pita string
	jumlah := 0
	i := 1

	for {
		fmt.Print("Bunga ", i, ": ")
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita = pita + bunga + " - "
		jumlah++
		i++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}

package main

import "fmt"

func main() {
	var N int
	var bunga string
	var pita string

	fmt.Print("N: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		fmt.Print("Bunga ", i, ": ")
		fmt.Scan(&bunga)
		pita = pita + bunga + " - "
	}

	fmt.Println("Pita:", pita)
}
