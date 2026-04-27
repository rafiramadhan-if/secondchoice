package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 {
		fmt.Println("Input tidak valid!")
		return
	}

	var ikan [1000]float64
	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	fmt.Println()

	var i int = 0
	var jumlahWadah int = 0
	var beratWadah float64 = 0

	for i < x {
		var total float64 = 0

		for j := 0; j < y && i < x; j++ {
			total += ikan[i]
			i++
		}

		jumlahWadah++
		beratWadah += total
		fmt.Printf("Wadah %d: %.2f\n", jumlahWadah, total)
	}

	fmt.Println()
	rataRata := beratWadah / float64(jumlahWadah)
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", rataRata)
}
