latihan 1
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
latihan 2
package main

import "fmt"

func main() {
	var tahun int
	var kabisat bool

	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)

	kabisat = (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	fmt.Println("Kabisat:", kabisat)
}
latihan 3
package main

import "fmt"

func main() {
	var r int
	const pi = 3.1415926535

	fmt.Print("Jari-jari: ")
	fmt.Scanln(&r)

	volume := (4.0 / 3.0) * pi * float64(r*r*r)
	luas := 4 * pi * float64(r*r)

	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
latihan 4
package main

import "fmt"

func main() {
	var celsius int

	fmt.Print("Masukkan Celsius: ")
	fmt.Scan(&celsius)

	reamur := celsius * 4 / 5
	fahrenheit := (celsius*9/5) + 32
	kelvin := celsius + 273

	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", kelvin)
}
latihan 5
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
latihan 6
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
latihan 7
package main

import "fmt"

func main() {
	var bunga string
	var pita string
	jumlah := 0
	i := 1

	for {
		fmt.Print("Bunga ", i, ": ")
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita = pita + bunga + " - "
		jumlah++
		i++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}
package main

import "fmt"

func main() {
	var N int
	var bunga string
	var pita string

	fmt.Print("N: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		fmt.Print("Bunga ", i, ": ")
		fmt.Scan(&bunga)
		pita = pita + bunga + " - "
	}

	fmt.Println("Pita:", pita)
}
latihan 8
package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var kiri, kanan float64
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)

		if kiri < 0 || kanan < 0 || (kiri+kanan) > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := math.Abs(kiri - kanan)
		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
latihan 9
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
latihan 10
package main

import "fmt"


func main() {
	var totalGram, kg, sisaGram, biayaKg, biayaSisa, totalBiaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&totalGram)

	kg = totalGram / 1000
	sisaGram = totalGram % 1000

	biayaKg = kg * 10000

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5
		} else {
			biayaSisa = sisaGram * 15
		}
	}

	totalBiaya = biayaKg + biayaSisa

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
} 
latihan 11
package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	fmt.Println("Nilai mata kuliah:", nmk)
}
latihan 12
package main

import "fmt"

func main() {
	var b, jumlahFaktor int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}
	fmt.Println()

	fmt.Printf("Prima: %t\n", jumlahFaktor == 2)
}
