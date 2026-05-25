package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id       string
	judul    string
	penulis  string
	penerbit string
	eksemplar int
	tahun    int
	rating   int
}

type DaftarBuku [7919]Buku

var Pustaka DaftarBuku
var nPustaka int

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	var i int
	i = 0
	for i < *n {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
		i = i + 1
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var idxMax, i int
	idxMax = 0
	i = 1
	for i < n {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
		i = i + 1
	}
	fmt.Println("Buku Terfavorit")
	fmt.Println("Judul    :", pustaka[idxMax].judul)
	fmt.Println("Penulis  :", pustaka[idxMax].penulis)
	fmt.Println("Penerbit :", pustaka[idxMax].penerbit)
	fmt.Println("Tahun    :", pustaka[idxMax].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j = j - 1
		}
		pustaka[j] = temp
		i = i + 1
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Buku Rating Tertinggi")
	batas := 5
	if n < 5 {
		batas = n
	}
	var i int
	i = 0
	for i < batas {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
		i = i + 1
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	var kiri, kanan, tengah int
	var ketemu bool
	kiri = 0
	kanan = n - 1
	ketemu = false

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			ketemu = true
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	fmt.Println("Hasil Pencarian Buku")
	if ketemu {
		fmt.Println("Judul     :", pustaka[tengah].judul)
		fmt.Println("Penulis   :", pustaka[tengah].penulis)
		fmt.Println("Penerbit  :", pustaka[tengah].penerbit)
		fmt.Println("Tahun     :", pustaka[tengah].tahun)
		fmt.Println("Eksemplar :", pustaka[tengah].eksemplar)
		fmt.Println("Rating    :", pustaka[tengah].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var ratingCari int

	DaftarkanBuku(&Pustaka, &nPustaka)

	CetakTerfavorit(Pustaka, nPustaka)

	UrutBuku(&Pustaka, nPustaka)

	Cetak5Terbaru(Pustaka, nPustaka)

	fmt.Scan(&ratingCari)
	CariBuku(Pustaka, nPustaka, ratingCari)
}
