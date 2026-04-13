package main

import "fmt"

func cetakPola(n int) {
	if n == 1 {
		fmt.Print(n)
		return
	}
	fmt.Print(n, " ")
	cetakPola(n - 1)
	fmt.Print(" ", n)
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakPola(n)
	fmt.Println()
}
