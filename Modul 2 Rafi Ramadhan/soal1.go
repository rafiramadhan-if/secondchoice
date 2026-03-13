package main

import "fmt"

func faktorial(n int) int {
	var hasil int
	var i int

	hasil = 1
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutation(n int, r int) int {
	var hasil int
	hasil = faktorial(n) / faktorial(n-r)
	return hasil
}

func combination(n int, r int) int {
	var hasil int
	hasil = faktorial(n) / (faktorial(r) * faktorial(n-r))
	return hasil
}

func main() {
	var a, b, c, d int
	var p1, k1, p2, k2 int

	fmt.Scan(&a, &b, &c, &d)

	p1 = permutation(a, c)
	k1 = combination(a, c)

	p2 = permutation(b, d)
	k2 = combination(b, d)

	fmt.Println(p1, k1)
	fmt.Println(p2, k2)
}
