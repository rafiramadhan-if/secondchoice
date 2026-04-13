package main

import "fmt"

func faktor(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan N: ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, ": ")
	faktor(n, 1)
	fmt.Println()
}
