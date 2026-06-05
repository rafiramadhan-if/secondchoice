package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var sum float64

	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	
	for i := 1; i <= n; i++ {
		suku := 1.0 / float64(2*i-1)

		if i%2 == 0 {
			suku = -suku
		}

		sum += suku
	}

	pi := 4 * sum
	fmt.Printf("Hasil PI (%d suku) : %.10f\n", n, pi)

	i := 1
	for {
		s1 := 1.0 / float64(2*i-1)
		s2 := 1.0 / float64(2*(i+1)-1)

		if i%2 == 0 {
			s1 = -s1
		}

		if (i+1)%2 == 0 {
			s2 = -s2
		}

		if math.Abs(s1-s2) <= 0.00001 {
			fmt.Printf("Suku ke-%d     : %.10f\n", i, s1)
			fmt.Printf("Suku ke-%d     : %.10f\n", i+1, s2)
			fmt.Println("Pada i ke      :", i)
			break
		}

		i++
	}
}
