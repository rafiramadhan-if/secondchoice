package main

import "fmt"

func cetakBintang(a int, n int) {
	if a > n {
		return
	}
	for i := 0; i < a; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	cetakBintang(a+1, n)
}

func main() {
	var n int
	fmt.Print()
	fmt.Scan(&n)
	cetakBintang(1, n)
}
