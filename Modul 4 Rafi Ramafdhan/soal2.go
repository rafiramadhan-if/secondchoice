package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	var waktu int

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		
		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, namaPemenang string
	var soalPeserta, skorPeserta int
	
	var maxSoal = -1
	var minSkor = 999999

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		hitungSkor(&soalPeserta, &skorPeserta)

		if soalPeserta > maxSoal || (soalPeserta == maxSoal && skorPeserta < minSkor) {
			maxSoal = soalPeserta
			minSkor = skorPeserta
			namaPemenang = nama
		}
	}

	fmt.Println(namaPemenang, maxSoal, minSkor)
}
