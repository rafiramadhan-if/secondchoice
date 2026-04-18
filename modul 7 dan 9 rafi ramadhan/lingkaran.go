package main

import (
	"fmt"
	"math"
)

type titik struct{ x, y int }
type lingkaran struct {
	pusat  titik
	radius int
}

func jarak(p, q titik) float64 {
	return math.Hypot(float64(p.x-q.x), float64(p.y-q.y))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
	var l1, l2 lingkaran
	var p titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&p.x, &p.y)

	switch {
	case didalam(l1, p) && didalam(l2, p):
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case didalam(l1, p):
		fmt.Println("Titik di dalam lingkaran 1")
	case didalam(l2, p):
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
