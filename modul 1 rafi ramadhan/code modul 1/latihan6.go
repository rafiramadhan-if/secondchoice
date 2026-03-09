package main

import "fmt"

func main() {
	var w1, w2, w3, w4 string
	var berhasil bool = true

	for i := 1; i <= 5; i++ {
		fmt.Scan(&w1, &w2, &w3, &w4)

		benar := w1 == "merah" && w2 == "kuning" && w3 == "hijau" && w4 == "ungu"
		berhasil = berhasil && benar
	}

	fmt.Println("BERHASIL:", berhasil)
}
