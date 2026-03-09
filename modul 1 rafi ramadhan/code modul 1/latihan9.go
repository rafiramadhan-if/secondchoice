package main

import (
	"fmt"
)

func main() {
	var kInput int
	
	fmt.Print("Nilai K = ")
	fmt.Scan(&kInput)

	akarDua := 1.0
	for k := 0; k <= kInput; k++ {
		akarDua *= float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", akarDua)
}
