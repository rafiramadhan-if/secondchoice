package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TIPE DATA DOMINO a

type Domino struct {
	Sisi1 int
	Sisi2 int
	Nilai int
	Balak bool
}

type Dominoes struct {
	Kartu  []Domino
	Jumlah int
}

// OPERASI DASAR DOMINO

func BuatDomino(s1, s2 int) Domino {
	return Domino{
		Sisi1: s1,
		Sisi2: s2,
		Nilai: s1 + s2,
		Balak: s1 == s2,
	}
}

func BuatSetDomino() Dominoes {

	var D Dominoes

	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			D.Kartu = append(D.Kartu, BuatDomino(i, j))
		}
	}

	D.Jumlah = len(D.Kartu)

	return D
}

func kocokKartu(D *Dominoes) {

	for i := range D.Kartu {

		j := rand.Intn(len(D.Kartu))

		D.Kartu[i], D.Kartu[j] =
			D.Kartu[j], D.Kartu[i]
	}
}

func ambilKartu(D *Dominoes) Domino {

	if D.Jumlah == 0 {
		return Domino{}
	}

	D.Jumlah--

	kartu := D.Kartu[D.Jumlah]

	D.Kartu = D.Kartu[:D.Jumlah]

	return kartu
}

func gambarKartu(d Domino, suit int) int {

	if suit == 1 {
		return d.Sisi1
	}

	return d.Sisi2
}

func nilaiKartu(d Domino) int {
	return d.Nilai
}

// SOAL NOMOR 2

func galiKartu(D *Dominoes, target Domino) Domino {

	for D.Jumlah > 0 {

		kartu := ambilKartu(D)

		if kartu.Sisi1 == target.Sisi1 ||
			kartu.Sisi2 == target.Sisi1 ||
			kartu.Sisi1 == target.Sisi2 ||
			kartu.Sisi2 == target.Sisi2 {

			return kartu
		}
	}

	return Domino{}
}

func sepasangKartu(d1, d2 Domino) bool {

	total :=
		nilaiKartu(d1) +
			nilaiKartu(d2)

	return total == 12
}

// MESIN KARAKTER

var pita string
var idx int
var currentChar byte

func start(input string) {

	pita = input

	idx = 0

	if len(pita) > 0 {
		currentChar = pita[idx]
	}
}

func maju() {

	idx++

	if idx < len(pita) {
		currentChar = pita[idx]
	}
}

func eop() bool {
	return currentChar == '.'
}

func cc() byte {
	return currentChar
}

// OPERASI MESIN KARAKTER

func bacaSemuaKarakter(input string) {

	start(input)

	fmt.Print("Karakter terbaca : ")

	for !eop() {
		fmt.Printf("%c", cc())
		maju()
	}

	fmt.Println()
}

func hitungKarakter(input string) int {

	start(input)

	jumlah := 0

	for !eop() {
		jumlah++
		maju()
	}

	return jumlah
}

func hitungA(input string) int {

	start(input)

	jumlahA := 0

	for !eop() {

		if cc() == 'A' {
			jumlahA++
		}

		maju()
	}

	return jumlahA
}

func frekuensiA(input string) float64 {

	total := hitungKarakter(input)
	jumlahA := hitungA(input)

	if total == 0 {
		return 0
	}

	return float64(jumlahA) / float64(total)
}

func hitungLE(input string) int {

	start(input)

	jumlah := 0
	var prev byte = ' '

	for !eop() {

		if prev == 'L' && cc() == 'E' {
			jumlah++
		}

		prev = cc()

		maju()
	}

	return jumlah
}

// MAIN PROGRAM

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("===================================")
	fmt.Println("MESIN ABSTRAK DOMINO")
	fmt.Println("===================================")

	// Membuat set domino
	setDomino := BuatSetDomino()

	fmt.Println("Jumlah kartu awal :", setDomino.Jumlah)

	// Mengocok kartu
	kocokKartu(&setDomino)

	// Mengambil kartu
	kartu1 := ambilKartu(&setDomino)

	fmt.Println("\nKartu pertama:")
	fmt.Printf("(%d,%d)\n",
		kartu1.Sisi1,
		kartu1.Sisi2)

	fmt.Println("Nilai :", nilaiKartu(kartu1))
	fmt.Println("Balak :", kartu1.Balak)

	// gambarKartu
	fmt.Println("Suit 1 :", gambarKartu(kartu1, 1))
	fmt.Println("Suit 2 :", gambarKartu(kartu1, 2))

	// Ambil kartu kedua
	kartu2 := ambilKartu(&setDomino)

	fmt.Println("\nKartu kedua:")
	fmt.Printf("(%d,%d)\n",
		kartu2.Sisi1,
		kartu2.Sisi2)

	fmt.Println("Nilai :", nilaiKartu(kartu2))

	// sepasangKartu
	fmt.Println("\nApakah total kedua kartu = 12 ?")
	fmt.Println(sepasangKartu(kartu1, kartu2))

	// galiKartu
	fmt.Println("\nMencari kartu yang memiliki suit sama...")

	hasil := galiKartu(&setDomino, kartu1)

	fmt.Printf("Ditemukan kartu (%d,%d)\n",
		hasil.Sisi1,
		hasil.Sisi2)


	// MESIN KARAKTER
	

	fmt.Println("\n===================================")
	fmt.Println("MESIN ABSTRAK KARAKTER")
	fmt.Println("===================================")

	var input string

	fmt.Print("Masukkan kalimat (akhiri dengan titik): ")
	fmt.Scanln(&input)

	bacaSemuaKarakter(input)

	fmt.Println("Jumlah karakter :", hitungKarakter(input))

	fmt.Println("Jumlah huruf A :", hitungA(input))

	fmt.Printf("Frekuensi A : %.2f%%\n",
		frekuensiA(input)*100)

	fmt.Println("Jumlah pasangan LE :",
		hitungLE(input))
}
