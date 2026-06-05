package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	var inside int

	fmt.Print("Banyak Topping: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()

		
		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= 0.25 {
			inside++
		}
	}

	
	pi := 4.0 * float64(inside) / float64(n)

	fmt.Println("Topping pada Pizza:", inside)
	fmt.Printf("PI : %.10f\n", pi)
}
