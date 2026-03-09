package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var x, y, z string

	fmt.Scan(&a, &b, &c, &d, &e)
	fmt.Scan(&x, &y, &z)

	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	fmt.Printf("%c%c%c\n", x[0]+1, y[0]+1, z[0]+1)
}
